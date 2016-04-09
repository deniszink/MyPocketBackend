package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"backend/settings"
	"backend/routers"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}

