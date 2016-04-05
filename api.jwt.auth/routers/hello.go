package routers

import (

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"backend/api.jwt.auth/core/authentication"
	"backend/api.jwt.auth/controllers"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
