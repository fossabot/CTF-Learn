package config

import (
	"github.com/spf13/viper"
)

//type StatusCode struct{
//	Success int
//	LoginUserError int
//	UserExisted	int
//}
var (
	Success         int
	LoginUserError  int
	UserExisted     int
	NoUserExisted   int
	NilToken        int
	NoRole          int
	DeleteUserError int
	PasswordError   int
)

func init() {
	Config()
}
func Config() {
	v := viper.New()
	v.SetConfigName("Code")
	v.AddConfigPath("./config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Success = v.GetInt(`StatusCode.Success`)
	LoginUserError = v.GetInt("StatusCode.LoginUserError")
	UserExisted = v.GetInt("StatusCode.UserExisted")
	NoUserExisted = v.GetInt("StatusCode.NoUserExisted")
	NilToken = v.GetInt("StatusCode.NilToken")
	NoRole = v.GetInt("StatusCode.NoRole")
	DeleteUserError = v.GetInt("StatusCode.DeleteUserError")
	PasswordError = v.GetInt("StatusCode.PasswordError")
}
