package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepostiory(DB *gorm.DB) UserRepository {
	DB.AutoMigrate(&User{})
	return UserRepository{DB: DB}
}

func (u *UserRepository) SignUp(login, password string) bool {
	var user User
	if u.DB.Where("login = ?", login).Take(&user).Error == nil {
		return false
	} else {
		user.Login = login
		user.Password = password
		u.DB.Create(&user)
		return true
	}
}

func (u *UserRepository) SignIn(login, password string) (int, error) {
	var user User
	err := u.DB.Where("login = ? AND password = ?", login, password).Take(&user).Error
	return user.ID, err
}

func (u *UserRepository) FindUser(login string) (int, string, error) {
	var user User
	err := u.DB.Where("login = ?", login).Take(&user).Error
	return user.ID, user.Login, err
}

func (u *UserRepository) ChangePassword(id int, password string) {
	user := User{ID: id}
	u.DB.Model(&user).Update("password", password)
}

func (u *UserRepository) ChangeLogin(id int, login string) bool {
	var user User
	if u.DB.Where("login = ?", login).Take(&user).Error == nil {
		return false
	}
	user.ID = id
	u.DB.Model(&user).Update("login", login)
	return true
}

func (u *UserRepository) ChangeAvatar(id int, avatar string) {
	user := User{ID: id}
	u.DB.Model(&user).Update("Avatar", avatar)
}

func (u *UserRepository) FindUserById(id int) (int, string, string, error) {
	var user User
	err := u.DB.Where("id = ?", id).Take(&user).Error
	return user.ID, user.Login, user.Avatar, err
}

/*func (u *UserRepository) ChangeAvatar(login string) {

}*/
