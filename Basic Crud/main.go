package main

import (
	"Crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetUniq).Methods(http.MethodGet)
	router.HandleFunc("/users", server.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/users", server.Update).Methods(http.MethodPut)
	router.HandleFunc("/users", server.Delete).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5000", router))
	fmt.Println("Listen on port :5000")
}
