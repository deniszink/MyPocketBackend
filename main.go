package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"backend/settings"
	"backend/routers"
	"os"
	"log"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), n))
}

