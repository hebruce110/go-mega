package main

import (
	"go-mega/lesson5/controller"
	"go-mega/lesson5/model"
	"log"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	controller.Startup()
	log.Println("Open http://127.0.0.1:9999")
	http.ListenAndServe(":9999", nil)
}
