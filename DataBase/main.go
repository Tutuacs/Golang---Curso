package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "golang:golang@/devbook?"
	db, err := sql.Open("mysql", urlConexao)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão Aberta!")

	lines, err := db.Query("select * from Users")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()

	fmt.Println(lines)
}
