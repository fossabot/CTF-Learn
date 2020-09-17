package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

func Init() {
	var err error
	dsn := "learnlogin:012359clown@tcp(127.0.0.1:3306)/learnlogin?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	} else {
		fmt.Print("数据库连接成功！")
	}
	Db.AutoMigrate(&User{})
	//设置最大闲置链接数
	Db.DB().SetMaxIdleConns(10)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	Db.DB().SetConnMaxLifetime(10 * time.Second)
}
