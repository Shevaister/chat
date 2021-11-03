package user

import (
	um "go/pkg/models/user" //um = userModel

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepostiory(DB *gorm.DB) UserRepository {
	DB.AutoMigrate(&um.User{})
	return UserRepository{DB: DB}
}

func (u *UserRepository) SignUp(login, password string) bool {
	var user um.User
	if u.DB.Where("login = ?", login).Take(&user).Error == nil {
		return false
	} else {
		user.Login = login
		user.Password = password
		u.DB.Create(&user)
		return true
	}
}

func (u *UserRepository) SignIn(login string) (int, string, error) {
	var user um.User
	err := u.DB.Where("login = ?", login).Take(&user).Error
	return user.ID, user.Password, err
}

// user search
func (u *UserRepository) FindUser(login string) []um.User {
	var user []um.User
	u.DB.Where("login LIKE ? LIMIT 20", "%"+login+"%").Find(&user)
	return user
}

func (u *UserRepository) ChangePassword(id int, password string) {
	user := um.User{ID: id}
	u.DB.Model(&user).Update("password", password)
}

func (u *UserRepository) ChangeLogin(id int, login string) bool {
	var user um.User
	if u.DB.Where("login = ?", login).Take(&user).Error == nil {
		return false
	}
	user.ID = id
	u.DB.Model(&user).Update("login", login)
	return true
}

func (u *UserRepository) ChangeAvatar(id int, avatar string) {
	user := um.User{ID: id}
	u.DB.Model(&user).Update("Avatar", avatar)
}

func (u *UserRepository) FindUserById(id int) (int, string, string, error) {
	var user um.User
	err := u.DB.Where("id = ?", id).Take(&user).Error
	return user.ID, user.Login, user.Avatar, err
}
