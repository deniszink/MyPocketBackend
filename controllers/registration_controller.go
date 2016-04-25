package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"backend/services"
	"fmt"
	"io/ioutil"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	newUser := new(models.User)
	decode := json.NewDecoder(r.Body)
	decode.Decode(&newUser)

	fmt.Println(newUser)
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("body = %s", string(data))

	responseStatus, body := services.Registration(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(body)
}