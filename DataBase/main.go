package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "golang:golang@/devbook?chartset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão Aberta!")

	lines, err := db.Query("select * from usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()
}
