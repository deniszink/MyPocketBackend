package models

import "gopkg.in/mgo.v2-unstable/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}