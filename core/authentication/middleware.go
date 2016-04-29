package authentication

import (
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"encoding/json"
	"backend/models"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc){
	authBackend := InitJWTAuthenticationBackend()

	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	fmt.Println("middleware",authBackend.IsInBlackList(req.Header.Get("Authorization")))
	//fmt.Println("token is Valid", &token.Valid)
	fmt.Println("error is",err)

	if err == nil && token.Valid && !authBackend.IsInBlackList(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		response,_ := json.Marshal(&models.Error{
			Error: "No token in request",
		})
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write(response)
	}
}