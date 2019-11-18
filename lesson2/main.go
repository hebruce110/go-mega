package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
}

// IndexViewModel struct
type IndexViewModel struct {
	Title string
	User  User
	Posts []Post
}

// Post struct
type Post struct {
	User User
	Body string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "Bruce"}
		v := IndexViewModel{
			Title: "Homepage",
			User:  user,
			Posts: []Post{
				Post{
				User: User{Username:"Bonfy"},
				Body: "Beautiful day in Portland!",},
				Post{
					User: User{Username:"Rene"},
					Body: "The Avengers movie was so cool!",
				},},
		}
		//tpl, _ := template.New("").Parse(`<html>
        //    <head>
        //        <title>Home Page - Bonfy</title>
        //    </head>
        //    <body>
        //        <h1>Hello, {{.Username}}!</h1>
		//		<br>
		//		<div>{{.BlogTitle}}</div>
        //    </body>
        //</html>`)
        t,err := template.ParseFiles("templates/index.html")
        if err!=nil{
        	log.Panic(err)
		}
		t.Execute(w, &v)
	})
	log.Println("Open http://127.0.0.1:9999")
	http.ListenAndServe(":9999", nil)
}
