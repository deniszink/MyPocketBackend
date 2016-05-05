package routers

import (
	"github.com/gorilla/mux"
	"backend/controllers"
	"github.com/codegangsta/negroni"
	"backend/core/authentication"
)

func SetWalletsRoutes(router *mux.Router) *mux.Router {
	router.Handle("/users/{userId}/wallets", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.CreateWallet),
	)).Methods("POST")


	router.Handle("/users/{userId}/wallets",negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.GetAllWallets),
	)).Methods("GET")

	return router
}
