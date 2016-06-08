package main


import (
	"net/http"
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
	//http.ListenAndServe(":5000", n)

}

