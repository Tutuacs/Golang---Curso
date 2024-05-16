package db

import (
	"database/sql"
	"log"
)

// Open Database Connection
func Conect() (db *sql.DB, err error) {
	urlConexao := "golang:golang@/devbook?"
	db, err = sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal(err)
	}

	return
}
