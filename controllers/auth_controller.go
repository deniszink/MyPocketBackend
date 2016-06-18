package controllers

import (
	"net/http"
	"encoding/json"
	"backend/models"
	"backend/services"
	"fmt"
)

func Login(w http.ResponseWriter, r *http.Request) {

	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	if requestUser.Email == "" || requestUser.Password == "" {
		fmt.Println(requestUser)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print("Error not nil")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		response, _ := json.Marshal(&models.Message{
			Message: "Logout successful",
		})
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}