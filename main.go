package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"backend/settings"
	"backend/routers"
	"os"

)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), n)
	if err != nil {
		panic(err)
	}
}

