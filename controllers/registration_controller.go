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

	responseStatus, body := services.Registr(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(body)
}