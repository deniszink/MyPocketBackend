package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"backend/services"
	"github.com/gorilla/mux"
)

func CreateWallet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	wallet := new(models.Wallet)
	decode := json.NewDecoder(r.Body)
	decode.Decode(&wallet)

	responseStatus, body := services.CreateWallet(wallet)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(body)
}

func GetAllWallets(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	vars := mux.Vars(r)
	userId := vars["userId"]

	responseStatus, body := services.GetAllWalletsByUser(userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(body)
}
