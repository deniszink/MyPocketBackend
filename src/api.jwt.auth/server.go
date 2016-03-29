package main

import (
	"api.jwt.auth/settings"
	"api.jwt.auth/routers"
	"github.com/codegangsta/negroni"
	"net/http"

	"os"
	"log"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":"+port,n)
}

