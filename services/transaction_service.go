package services

import (
	"backend/models"
	"backend/core/store"
	"fmt"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func CreateTransaction(transaction *models.Transaction) (int, []byte) {
	mongo := store.ConnectMongo()

	isValid, message := isTransactionValid(transaction)
	if !isValid {
		response, _ := json.Marshal(&models.Error{
			Error: message,
		})
		return http.StatusBadRequest, response
	}

	if err := mongo.WriteDataTo(store.TableTransactions, transaction); err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(&models.Error{
			Error: "Something went wrong, can't create new transaction.",
		})
		return http.StatusInternalServerError, response
	} else {
		return http.StatusCreated, []byte("")
	}

}

func isTransactionValid(transaction *models.Transaction) (bool, string) {
	mongo := store.ConnectMongo()

	if isValid := (transaction.Amount > 0 && transaction.Amount != 0) &&
	(transaction.TransactionType >= 0 && transaction.TransactionType <= 1); !isValid {
		return false, "amount should be > 0, type should be 0 or 1"
	}

	wallet := new(models.Wallet)
	err := mongo.GetOne(store.TableWallets, bson.M{"_id":transaction.WalletId}, wallet)
	if err != nil {
		fmt.Println(err)
		return false, "Wallet doesn't exists"
	}

	if isUserExists, _ := mongo.IsExists(store.TableUsers, bson.M{"_id":wallet.UserID}); !isUserExists {
		return false, "User doesn't exists"
	}

	return true, ""

}
