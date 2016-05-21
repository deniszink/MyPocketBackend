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

	http.ListenAndServe(":5000", n)
	/*connect := redis.Connect()
	connect.SetValue("token","eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjIxNzkxMzcsImlhdCI6MTQ2MTkxOTkzNywic3ViIjoiIn0.JMNC18eVg9ZRbIywzZIc4jxSlrxJhnD-R9gPI_B48EGdFaOVuJm4b0Uk6LOHHjsvIh1NaJI39l1v94gih0cQcZP7j8OPdtrIt_ZDxSY499XY5dbgZMS1ftg-bk0VsP9YrAYZjhZ-8yesQlT9qpFliaGPHovdltgI9oFPURmIOOtumZTEPary-1X_Kt78vrfoz1TNEAwOuFXtbSaozFgVFiC_o6qsMqcS6fhfw9aOITa3yerQeVgKwKOk1F80i1zE7LacsYaPxqWMz9CR9cSHxuRVlsl1_nWEP_bZW8krP6wxg70x-L1Kze1yCO8SQ9ZTvak7DaublTwxOY83-hlOAQ")
	fmt.Println(connect.GetValue("token"))*/
}

