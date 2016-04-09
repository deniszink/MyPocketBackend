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
	port := os.Getenv("PORT")
	log.Println(port)
	err := http.ListenAndServe(":" + port, n)
	if err != nil {
		panic(err)
	}
}

