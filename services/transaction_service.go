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

type GET_Transactions struct {
	Transactions []models.Transaction `json:"transactions,omitempty" form:"transactions"`
}

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

func GetAllTransactions(walletId bson.ObjectId) (int,[]byte){
	mongo := store.ConnectMongo()

	if isWalletExists, err := mongo.IsExists(store.TableWallets,bson.M{"_id":walletId}); !isWalletExists || err != nil{
		fmt.Println(err)
		response,_ := json.Marshal(&models.Error{
			Error: "Wallet doesn't exists",
		})
		return http.StatusBadRequest,response
	}

	var transactions []models.Transaction
	if err := mongo.FindAll(store.TableTransactions,bson.M{"walletId":walletId},&transactions); err != nil{
		fmt.Println(err)
		return http.StatusInternalServerError, []byte("")
	}

	response, _ := json.Marshal(&GET_Transactions{
		Transactions: transactions,
	})

	//response, _ := json.Marshal(transactions)

	return http.StatusOK, response
}
