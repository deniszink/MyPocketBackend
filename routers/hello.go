package routers

import (

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"backend/core/authentication"
	"backend/controllers"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
