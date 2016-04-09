package routers

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"log"
	"backend/core/authentication"
	"backend/controllers"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router{
	log.Print("SetAuthenticationRoutes")
	router.HandleFunc("/token-auth",controllers.Login).Methods("POST")

	router.Handle("/refresh-token-auth",negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")

	router.Handle("/logout",negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Logout),
	)).Methods("GET")

	return router
}