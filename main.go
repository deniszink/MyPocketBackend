package main

import (
	"backend/models"
	"net/http"
	"os"
	"backend/settings"
	"backend/routers"
	"github.com/codegangsta/negroni"
)

type Model struct {
	 *models.User `json:"user"`
	 *models.Token
}
func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":"+os.Getenv("PORT"), n)
}

