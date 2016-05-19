package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"id,omitempty" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
