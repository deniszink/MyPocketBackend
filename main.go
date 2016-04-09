package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"backend/api.jwt.auth/settings"
	"backend/api.jwt.auth/routers"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}

