package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"backend/core/authentication"
	"backend/controllers"
)

func SetTransactionRoutes(router *mux.Router) *mux.Router {
	router.Handle("/wallets/{walletId}/transactions", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.CreateTransaction),
	)).Methods("POST")

	router.Handle("/wallets/{walletId}/transactions", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.GetAllIncomeTransactions),
	)).Methods("GET").Queries("filter", "income")

	router.Handle("/wallets/{walletId}/transactions", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.GetAllExpenseTransactions),
	)).Methods("GET").Queries("filter", "expense")

	router.Handle("/wallets/{walletId}/transactions", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.GetAllTransactions),
	)).Methods("GET")

	return router
}
