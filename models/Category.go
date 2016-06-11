package models

import "gopkg.in/mgo.v2/bson"

type Category struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}
