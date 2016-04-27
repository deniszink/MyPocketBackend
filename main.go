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

	http.ListenAndServe(":"+os.Getenv("PORT"), n)
}

