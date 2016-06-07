package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"backend/services"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	newUser := new(models.User)
	decode := json.NewDecoder(r.Body)
	decode.Decode(&newUser)

	if newUser.Email == "" || newUser.Password == "" || newUser.Username == "" {
		body, _ := json.Marshal(&models.Error{
			Error: "Body is not valid, check your request",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
	} else {
		responseStatus, body := services.Registration(newUser)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		w.Write(body)
	}
}