package main


import (
	"net/http"
	"os"
	"backend/settings"
	"backend/routers"
	"github.com/codegangsta/negroni"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":"+os.Getenv("PORT"), n)
}

