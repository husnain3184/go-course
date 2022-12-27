package main

import (
	"fmt"
	handler "golang/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/signup", handler.Signup)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.HandleFunc("/", handler.Static)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// http.Handle("/css/", http.FileServer(http.Dir("../static")))

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
