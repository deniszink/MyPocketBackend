package main

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"backend/settings"
	"backend/routers"
	"backend/services"
	"os"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	services.CreateCategoies()
	http.ListenAndServe(":"+os.Getenv("PORT"), n)
	//http.ListenAndServe(":5000", n)


}

