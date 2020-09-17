package main

import (
	"LearnLogin/database"
	"LearnLogin/router"
)

func main() {
	database.Init()
	router.Router()
	defer database.Db.Close()
}
