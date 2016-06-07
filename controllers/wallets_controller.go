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

	if !validateWallet(wallet){
		body,_ := json.Marshal(&models.Error{
			Error: "Body is not valid check your request",
		})
		WriteResponse(w,http.StatusBadRequest,body)
	}else {

		if isObjectIdHex := bson.IsObjectIdHex(userId); !isObjectIdHex {
			body, _ := json.Marshal(&models.Error{
				Error: "Invalid userId format, must be 12 length, check your input data",
			})
			WriteResponse(w, http.StatusBadRequest, body)
		} else {
			wallet.UserID = bson.ObjectIdHex(userId)
			responseStatus, body := services.CreateWallet(wallet)
			WriteResponse(w, responseStatus, body)
		}
	}
}

func GetAllWallets(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	responseStatus, body := services.GetAllWalletsByUser(userId)
	WriteResponse(w, responseStatus, body)
}

func validateWallet(wallet *models.Wallet) bool {
	return wallet.WalletName != "" || wallet.Balance != 0 || wallet.Currency != ""
}
