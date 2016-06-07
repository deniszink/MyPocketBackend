package services

import (
	"backend/models"
	"backend/core/store"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"net/http"
	"fmt"
)

func CreateWallet(wallet *models.Wallet) (int, []byte) {
	mongo := store.ConnectMongo()

	//check if user exists
	if isUserExists, _ := mongo.IsExists(store.TableUsers, bson.M{"_id":wallet.UserID}); isUserExists {
		//if so check if wallet already exists
		if isWalletExists, _ := mongo.IsExists(store.TableWallets, bson.M{"walletname":wallet.WalletName, "userId":wallet.UserID}); isWalletExists {
			data, _ := json.Marshal(&models.Error{
				Error: "User can't has two equal wallet",
			})
			return http.StatusBadRequest, data
		} else {
			if err := mongo.WriteDataTo(store.TableWallets, wallet); err != nil {
				panic(err)
			}

			wallet := new(models.Wallet)
			err := mongo.GetOne(store.TableWallets, bson.M{"userId":wallet.UserID, "walletName":wallet.WalletName}, wallet)
			if err != nil{
				panic(err)
			}
			response, _ := json.Marshal(wallet)
			return http.StatusCreated, response
		}
	} else {
		response, _ := json.Marshal(&models.Error{
			Error: "Can't find user with this ID",
		})
		return http.StatusBadRequest, response
	}
}

func GetAllWalletsByUser(userID string) (int, []byte) {
	var wallets []models.Wallet

	mongo := store.ConnectMongo()
	if err := mongo.FindAll(store.TableWallets, bson.M{"userId":bson.ObjectIdHex(userID)}, &wallets); err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(&models.Error{
			Error: "Error while trying get all wallet by userId",
		})
		return http.StatusInternalServerError, []byte(response)
	}

	if (len(wallets) == 0) {
		return http.StatusOK, []byte("[]")
	} else {
		data, _ := json.Marshal(wallets)
		return http.StatusOK, []byte(data)
	}
}