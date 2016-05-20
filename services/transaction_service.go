package services

import (
	"backend/models"
	"backend/core/store"
	"fmt"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"strings"
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
		c := make(chan uint8)
		go doTransaction(transaction, c)
		if res := <-c; res == 1 {
			return http.StatusCreated, []byte("")
		} else {
			return http.StatusInternalServerError, []byte("")
		}
	}

}

func isTransactionValid(transaction *models.Transaction) (bool, string) {
	mongo := store.ConnectMongo()

	if isValid := (transaction.Amount > 0 && transaction.Amount != 0) &&
	(strings.EqualFold(transaction.TransactionType,"income") || strings.EqualFold(transaction.TransactionType, "expense")); !isValid {
		return false, "amount should be > 0, type should be 'expense' or 'income'"
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

func doTransaction(transaction *models.Transaction, c chan <- uint8) error {
	mongo := store.ConnectMongo()
	amount := transaction.Amount

	wallet := new(models.Wallet)
	mongo.GetOne(store.TableWallets, bson.M{"_id": transaction.WalletId}, wallet)
	if (strings.EqualFold("expense",transaction.TransactionType)) {
		wallet.Balance -= amount
	} else {
		wallet.Balance += amount
	}

	err := mongo.Update(store.TableWallets, bson.M{"_id":transaction.WalletId}, wallet)
	if err != nil {
		c <- 0
	} else {
		c <- 1
	}
	return err
}
