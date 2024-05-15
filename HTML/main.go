package main

import (
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/users", usersFunc)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homeFunc(w http.ResponseWriter, _ *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
	w.Write([]byte("Olá Mundo!"))
}

func usersFunc(w http.ResponseWriter, _ *http.Request) {

	user := user{"Arthur Azevedo da Silva", "arthursilva.rs1@gmail.com"}
	templates.ExecuteTemplate(w, "user.html", user)
	w.Write([]byte("User Pages!"))
}
