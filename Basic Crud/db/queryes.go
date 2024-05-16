package db

import "log"

func Create() {
	db, err := Conect()
	if err != nil {
		log.Fatal(err)
	}

}
