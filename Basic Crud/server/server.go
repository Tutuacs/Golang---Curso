package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	id string `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	emaill string `json:"email"`
	age int `json:"age"`
} 

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := ioutil.ReadAll(r.Body){
		if err != nil {
			w.Write([]byte("Falha ao ler o corpo da requisição!"))
			return
		}

		var user User

		if err = json.Unmarshal(req, &user); err != nil {
			w.Write([]byte("Error while convert req to user!"))
			return
		}
	}
}

func GetUniq(w http.ResponseWriter, r *http.Request) {

}

func GetAll(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
