package services

import (
	"backend/models"
	"backend/core/store"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"net/http"

	"fmt"
)

func CreateWallet(wallet *models.Wallet) (int,[]byte){
	mongo := store.ConnectMongo()

	isExists,err := mongo.IsExists(store.TableWallets, bson.M{"walletname":wallet.WalletName,"userid":wallet.UserID})
	if !isExists{
		fmt.Println(err)
		if err := mongo.WriteDataTo(store.TableWallets,wallet); err != nil{
			panic(err)
		}
		response,_ := json.Marshal(&models.Message{"Wallet succesfully created"})
		return http.StatusCreated, response
	}

	data,_ := json.Marshal(&models.Error{
		Error: "User can't has two equal wallet",
	})
	return http.StatusBadRequest, data
}