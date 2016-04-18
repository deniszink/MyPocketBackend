package services

import (
	"backend/models"
	"backend/core/store"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"

)

func Registration(newUser *models.User)(int,[]byte){
	mongo := store.ConnectMongo()

	err := mongo.FindOne(store.TableUsers,bson.M{"email":newUser.Email}, nil); if err != nil {
		mongo.WriteDataTo(store.TableUsers,newUser)
		return http.StatusCreated, []byte("")
	}

	response, _ := json.Marshal(models.Error{
		Error: "User with this email already exists",
	})
	return http.StatusBadRequest, response

}
