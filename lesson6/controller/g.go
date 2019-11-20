package controller

import (
	"github.com/gorilla/sessions"
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store *sessions.CookieStore
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("xasdf234asdfasdfa531123"))
	sessionName = "go-mega"
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}