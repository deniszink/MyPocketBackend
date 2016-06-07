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
	isExists,_ := mongo.IsExists(store.TableUsers,bson.M{"email":newUser.Email});
	if !isExists {
		mongo.WriteDataTo(store.TableUsers,newUser)

		user := new(models.User)
		mongo.GetOne(store.TableUsers,bson.M{"email":newUser.Email},user)
		body,_ := json.Marshal(user)
		return http.StatusCreated, []byte(body)
	}

	response, _ := json.Marshal(models.Error{
		Error: "User with this email already exists",
	})
	return http.StatusBadRequest, []byte(response)
}
