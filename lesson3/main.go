package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string
}

// IndexViewModel struct
type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

// Post struct
type Post struct {
	User User
	Body string
}

func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _,fi := range fis {
		func(){
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}
			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}()
	}
	return result
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
        //t,err := template.ParseFiles("templates/index.html")
        t := PopulateTemplates()
        t["index.html"].Execute(w,&v)
	})
	log.Println("Open http://127.0.0.1:9999")
	http.ListenAndServe(":9999", nil)
}
