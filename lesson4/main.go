package main

import (
	"go-mega/lesson4/controller"
	"log"
	"net/http"
)

func main() {
	controller.Startup()
	log.Println("Open http://127.0.0.1:9999")
	http.ListenAndServe(":9999", nil)
}
