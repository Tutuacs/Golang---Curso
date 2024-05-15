package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/users", usersFunc)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homeFunc(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usersFunc(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("User Pages!"))
}
