package controllers

import (
	"net/http"
	"backend/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"backend/services"
	"fmt"
	"backend/core/store"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	transaction := new(models.Transaction)
	decode := json.NewDecoder(r.Body)
	decode.Decode(transaction)
	fmt.Println("TRANSACTION ",transaction)

	vars := mux.Vars(r)
	walletId := vars["walletId"]

	if !validateTransaction(transaction){
		response, _ := json.Marshal(&models.Error{
			Error: "Body in not valid, check your request",
		})
		WriteResponse(w,http.StatusBadRequest,response)
	}else {
		if doWalletIDValidation(w, walletId) {
			if isExist,_ := isWalletExists(bson.ObjectId(walletId)); isExist {
				validWalletId := bson.ObjectIdHex(walletId)
				transaction.WalletId = validWalletId
				fmt.Println(transaction)
				code, body := services.CreateTransaction(transaction)
				WriteResponse(w, code, body)
			}else{
				response,_ := json.Marshal(&models.Error{
					Error: "Wallet with this id doesn't exists",
				})
				WriteResponse(w,http.StatusBadRequest,response)
			}
		}
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
		WriteResponse(w,code,body)
	}
}

func GetAllIncomeTransactions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	fmt.Println(vars)

	if doWalletIDValidation(w, walletId) {
		validWalletId := bson.ObjectIdHex(walletId)
		code, body := services.GetAllIncomeTransaction(validWalletId)
		WriteResponse(w,code,body)
	}
}

func GetAllExpenseTransactions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]

	if doWalletIDValidation(w, walletId) {
		validWalletId := bson.ObjectIdHex(walletId)
		code, body := services.GetAllExpenseTransactions(validWalletId)
		WriteResponse(w,code,body)
	}
}

func doWalletIDValidation(w http.ResponseWriter, walletId string) bool {
	if hex := bson.IsObjectIdHex(walletId); !hex {
		response, _ := json.Marshal(&models.Error{
			Error: "Wallet id is not valid, please check your input data.",
		})
		WriteResponse(w,http.StatusBadRequest,response)
		return false
	}
	return true
}

func validateTransaction(t *models.Transaction) bool {
	return t.Amount != 0 || t.Type != "" || t.UnixDateTime != 0
}

func isWalletExists(walletId bson.ObjectId) (bool,error) {
	mongo := store.ConnectMongo()
	return mongo.IsExists(store.TableWallets,bson.M{"_id":walletId})
}


