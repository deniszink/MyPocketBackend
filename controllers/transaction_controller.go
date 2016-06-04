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

	if doWalletIDValidation(w, walletId) {
		validWalletId := bson.ObjectIdHex(walletId)
		transaction.WalletId = validWalletId
		fmt.Println(transaction)
		code, body := services.CreateTransaction(transaction)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(body)
	}

}

func GetAllTransactions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	fmt.Println(vars)
	fmt.Println(vars["filter"])

	if doWalletIDValidation(w, walletId) {
		validWalletId := bson.ObjectIdHex(walletId)
		code, body := services.GetAllTransactions(validWalletId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(body)
	}
}

func GetAllIncomeTransactions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	fmt.Println(vars)

	if doWalletIDValidation(w, walletId) {
		validWalletId := bson.ObjectIdHex(walletId)
		code, body := services.GetAllIncomeTransaction(validWalletId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(body)
	}
}

func GetAllExpenseTransactions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	if doWalletIDValidation(w, walletId) {
		validWalletId := bson.ObjectIdHex(walletId)
		code, body := services.GetAllExpenseTransactions(validWalletId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(body)
	}
}

func doWalletIDValidation(w http.ResponseWriter, walletId string) bool {
	if hex := bson.IsObjectIdHex(walletId); !hex {
		response, _ := json.Marshal(&models.Error{
			Error: "Wallet id is not valid, please check your input data.",
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return false
	}
	return true
}