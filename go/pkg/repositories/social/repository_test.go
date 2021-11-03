package social

import (
	sm "go/pkg/models/social"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	socialRepository := ProvideSocialRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `socials` WHERE user1_id = ? AND user2_id = ? and status = ? LIMIT 1").
		WithArgs(1, 2, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user1_id", "user2_id", "status"}))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `socials` (`user1_id`,`user2_id`,`status`) VALUES (?,?,?) ").WithArgs(1, 2, 1).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	success := socialRepository.AddUser(1, 2)

	require.Equal(t, success, true)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestBlockUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	socialRepository := ProvideSocialRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `socials` WHERE user1_id = ? AND user2_id = ? and status = ? LIMIT 1").
		WithArgs(1, 2, 0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user1_id", "user2_id", "status"}))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `socials` (`user1_id`,`user2_id`,`status`) VALUES (?,?,?) ").WithArgs(1, 2, 0).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	success := socialRepository.BlockUser(1, 2)

	require.Equal(t, success, true)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUnfriendUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	socialRepository := ProvideSocialRepostiory(gdb)
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `socials` WHERE user1_id = ? AND user2_id = ? AND status = 1").WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	socialRepository.UnfriendUser(1, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFriends(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	socialRepository := ProvideSocialRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `socials` WHERE user1_id = ? AND status = 1").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user1_id", "user2_id", "status"}).
				AddRow(1, 1, 2, 1))

	friends := socialRepository.GetFriends(1)

	require.Equal(t, []sm.Social{{ID: 1, User1ID: 1, User2ID: 2, Status: 1}}, friends)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCheckBlocked(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	socialRepository := ProvideSocialRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `socials` WHERE user1_id = ? AND user2_id = ? and status = ? LIMIT 1").
		WithArgs(1, 2, 0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "user1_id", "user2_id", "status"}).
				AddRow(1, 1, 2, 0))

	blocked := socialRepository.CheckBlocked(1, 2)

	require.Equal(t, blocked, true)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
