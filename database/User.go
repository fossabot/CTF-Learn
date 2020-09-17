package database

import (
	"LearnLogin/config"
	"LearnLogin/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Role     string `gorm:"not null" json:"role" binding:"required"`
}

//验证用户是否存在
func CheckUser(username string, role string) (code int) {
	var user User
	if role != "user" && role != "admin" {
		return config.NoRole
	}
	Db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return config.UserExisted
	}
	return config.Success
}

//用户登录验证
func CheckLogin(username string, password string, role string) (code int) {
	var user User
	Db.Where("username = ?", username).First(&user)
	//判断用户是否存在
	if user.ID == 0 {
		return config.NoUserExisted
	}
	if utils.EncryptionPwd(password) != user.Password {
		return config.PasswordError
	}
	if user.Role != role {
		return config.NoRole
	}
	return config.Success
}

//删除用户
func DeleteUser(username string) (int, error) {
	var user User
	err := Db.Where("username = ?", username).Delete(&user).Error
	if err != nil {
		return config.DeleteUserError, err
	}
	return config.Success, nil
}
