package services

import (
	"backend/models"
	"backend/core/store"
	"github.com/go-mgo/mgo/bson"
	"net/http"
)

func Registr(newUser *models.User)(int,[]byte){
	var existUser models.User
	mongo := store.ConnectMongo()
	mongo.FindOne(store.TableUsers,bson.M{"email":newUser.Email}, &existUser)

	if existUser != nil{
		return http.StatusBadRequest, []byte("User with this email already exist")
	}

	mongo.WriteDataTo(store.TableUsers,newUser)
	return http.StatusCreated, []byte("")
}
