package main

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"backend/settings"
	"backend/routers"
	"os"
	"backend/services"
)

func main() {
	services.CreateCategoies()
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), n)
	//http.ListenAndServe(":5000", n)
}

