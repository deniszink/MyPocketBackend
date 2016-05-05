package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"backend/services"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func CreateWallet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	wallet := new(models.Wallet)
	decode := json.NewDecoder(r.Body)
	decode.Decode(&wallet)

	vars := mux.Vars(r)
	userId := vars["userId"]

	if isObjectIdHex := bson.IsObjectIdHex(userId); !isObjectIdHex {
		body, _ := json.Marshal(&models.Error{
			Error: "Invalid userId format, must be 12 length, check your input data",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
	} else {
		wallet.UserID = bson.ObjectIdHex(userId)

		responseStatus, body := services.CreateWallet(wallet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		w.Write(body)
	}
}

func GetAllWallets(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	responseStatus, body := services.GetAllWalletsByUser(userId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(body)
}
