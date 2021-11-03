package message

import (
	mm "go/pkg/models/message"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetAllChatMessages(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	messageRepository := ProvideMessageRepostiory(gdb)

	rs := []mm.Message{{ID: 1, ChatID: 1, SenderID: 1, Text: "hi", CreatedAt: time.Now()}, {ID: 2, ChatID: 1, SenderID: 2, Text: "hi", CreatedAt: time.Now()}}

	mock.ExpectQuery(
		"SELECT * FROM `messages` WHERE chat_id = ?").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "chat_id", "sender_id", "text", "created_at"}).
				AddRow(rs[0].ID, rs[0].ChatID, rs[0].SenderID, rs[0].Text, rs[0].CreatedAt).
				AddRow(rs[1].ID, rs[1].ChatID, rs[1].SenderID, rs[1].Text, rs[1].CreatedAt))

	messages := messageRepository.GetAllChatMessages(1)

	require.Equal(t, rs, messages)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetLastChatMessage(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	messageRepository := ProvideMessageRepostiory(gdb)

	rs := mm.Message{ID: 1, ChatID: 1, SenderID: 1, Text: "hi", CreatedAt: time.Now()}

	mock.ExpectQuery(
		"SELECT * FROM `messages` WHERE chat_id = ? ORDER BY `messages`.`id` DESC LIMIT 1").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "chat_id", "sender_id", "text", "created_at"}).
				AddRow(rs.ID, rs.ChatID, rs.SenderID, rs.Text, rs.CreatedAt))

	message, err := messageRepository.GetLastChatMessage(1)

	require.NoError(t, err)
	require.Equal(t, rs, message)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
