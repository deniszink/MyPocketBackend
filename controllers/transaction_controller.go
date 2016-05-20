package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"backend/services"
	"fmt"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	transaction := new(models.Transaction)
	decode := json.NewDecoder(r.Body)
	decode.Decode(transaction)

	vars := mux.Vars(r)
	walletId := vars["walletId"]

	if hex := bson.IsObjectIdHex(walletId); !hex {
		response, _ := json.Marshal(&models.Error{
			Error: "Wallet id is not valid, please check your input data.",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	} else {

		fmt.Println("Transaction type :",transaction.TransactionType)

		validWalletId := bson.ObjectIdHex(walletId)
		transaction.WalletId = validWalletId
		code, body := services.CreateTransaction(transaction)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(body)
	}

}