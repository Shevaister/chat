package user

import (
	um "go/pkg/models/user"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSignUp(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `users` WHERE login = ? LIMIT 1").
		WithArgs("qwerty21").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "login", "password", "avatar"}))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`login`,`password`,`avatar`) VALUES (?,?,?) ").WithArgs("qwerty21", "qwerty21", "").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	success := userRepository.SignUp("qwerty21", "qwerty21")

	require.Equal(t, true, success)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSignIn(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `users` WHERE login = ? LIMIT 1").
		WithArgs("qwerty21").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "login", "password", "avatar"}).
				AddRow(1, "qwerty21", "qwerty21", ""))

	id, password, err := userRepository.SignIn("qwerty21")

	require.Equal(t, 1, id)
	require.Equal(t, "qwerty21", password)
	require.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := []um.User{{ID: 1, Login: "zwerty21", Password: "qwerty21", Avatar: ""}, {ID: 2, Login: "zwerty22", Password: "qwerty21", Avatar: ""}}

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `users` WHERE login LIKE ?").
		WithArgs("%zwerty%").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "login", "password", "avatar"}).
				AddRow(rs[0].ID, rs[0].Login, rs[0].Password, rs[0].Avatar).
				AddRow(rs[1].ID, rs[1].Login, rs[1].Password, rs[1].Avatar))

	users := userRepository.FindUser("zwerty")

	require.Equal(t, rs, users)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestChangePassword(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := map[string]interface{}{"id": 1, "password": "qwerty21"}

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `password`=? WHERE `id` = ?").WithArgs(rs["password"], rs["id"].(int)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	userRepository.ChangePassword(rs["id"].(int), rs["password"].(string))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestChangeLogin(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := map[string]interface{}{"id": 1, "login": "qwerty21"}

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `login`=? WHERE `id` = ?").WithArgs(rs["login"], rs["id"].(int)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	userRepository.ChangeLogin(rs["id"].(int), rs["login"].(string))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestChangeAvatar(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := map[string]interface{}{"id": 1, "avatar": "1.png"}

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `avatar`=? WHERE `id` = ?").WithArgs(rs["avatar"], rs["id"].(int)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	userRepository.ChangeAvatar(rs["id"].(int), rs["avatar"].(string))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindUserById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}))

	rs := um.User{ID: 1, Login: "qwerty21", Password: "qwerty21", Avatar: ""}

	userRepository := ProvideUserRepostiory(gdb)

	mock.ExpectQuery(
		"SELECT * FROM `users` WHERE id = ? LIMIT 1").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "login", "password", "avatar"}).
				AddRow(rs.ID, rs.Login, rs.Password, rs.Avatar))

	userId, userLogin, userAvatar, err := userRepository.FindUserById(rs.ID)

	require.NoError(t, err)
	require.Equal(t, rs.ID, userId)
	require.Equal(t, rs.Login, userLogin)
	require.Equal(t, rs.Avatar, userAvatar)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
