package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userId, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("New user inserted with ID %d", userId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
