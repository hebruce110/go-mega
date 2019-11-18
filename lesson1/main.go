package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// User struct
type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
		fmt.Fprint(w,"hello2\n")
		io.WriteString(w,"hello3\n")
	})
	log.Println("Open http://127.0.0.1:9999")
	http.ListenAndServe(":9999", nil)
}