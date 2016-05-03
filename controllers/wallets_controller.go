package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"fmt"
	"backend/services"
)

func CreateWallet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	wallet := new(models.Wallet)
	decode := json.NewDecoder(r.Body)
	decode.Decode(&wallet)

	fmt.Println(wallet)

	responseStatus, body := services.CreateWallet(wallet)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(body)
}
