package service

import (
	"LearnLogin/database"
	"LearnLogin/utils"
)

//加密用户密码并注册
func AddUser(data *database.User) {
	data.Password = utils.EncryptionPwd(data.Password)
	err := database.Db.Create(&data).Error
	if err != nil {
		panic(err)
	}
}
