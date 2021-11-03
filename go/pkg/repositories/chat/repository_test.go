package chat

import (
	cm "go/pkg/models/chat"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetChatID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	chatRepository := ProvideChatRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `chats` WHERE user1_id = ? AND user2_id = ? OR user1_id = ? AND user2_id = ? LIMIT 1").
		WithArgs(2, 8, 8, 2).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user1_id", "user2_id"}))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `chats` (`user1_id`,`user2_id`) VALUES (?,?)").WithArgs(2, 8).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	chatId := chatRepository.GetChatID(2, 8)

	require.Equal(t, 1, chatId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetChatsID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	chatRepository := ProvideChatRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `chats` WHERE user1_id = ? OR user2_id = ?").
		WithArgs(1, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user1_id", "user2_id"}).
				AddRow(1, 1, 2).
				AddRow(2, 1, 4))

	chatsId := chatRepository.GetChatsID(1)

	require.Equal(t, []cm.Chat{{ID: 1, User1ID: 1, User2ID: 2}, {ID: 2, User1ID: 1, User2ID: 4}}, chatsId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
