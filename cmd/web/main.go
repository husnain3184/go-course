package main

import (
	"fmt"
	config "golang/pkg/config"
	handler "golang/pkg/handlers"
	render "golang/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/services", handler.Services)
	http.HandleFunc("/blog", handler.Blog)
	http.HandleFunc("/contact", handler.Contact)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/signup", handler.Signup)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.HandleFunc("/", handler.Static)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// http.Handle("/css/", http.FileServer(http.Dir("../static")))

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
