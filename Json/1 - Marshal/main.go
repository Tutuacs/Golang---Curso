package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	dog := cachorro{"Rex", "Dálmata", 3}

	fmt.Println(dog)

	cachorro, err := json.Marshal(dog)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorro)
	fmt.Println(bytes.NewBuffer(cachorro))

	dog2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorro2)
	fmt.Println(bytes.NewBuffer(cachorro2))
}
