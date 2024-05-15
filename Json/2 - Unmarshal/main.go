package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}
type cachorro2 struct {
	Nome  string `json:"-"` // Desta forma é possivel censurar algo ao passar ou receber dados em JSON
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	dogJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var dog cachorro

	if err := json.Unmarshal([]byte(dogJSON), &dog); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dog)

}
