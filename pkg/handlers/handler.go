package handler

import (
	"golang/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

// func Static(w http.ResponseWriter, r *http.Request) {
// 	http.FileServer(http.Dir("css/"))

// }

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// Login is the handler for the about page
func Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl")
}

// signup is the handler for the about page
func Signup(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "signup.page.tmpl")
}
