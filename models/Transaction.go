package models

import "gopkg.in/mgo.v2/bson"

type Transaction struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	WalletId     bson.ObjectId `bson:"walletId,omitempty" json:"walletId"`
	Amount       float32 `json:"amount" form:"amount"`
	Type         string`json:"type" bson:"type"`
	Category     string `json:"categoryId" bson:"categoryId"`
	UnixDateTime int64 `json:"unixDataTime" bson:"unixDataTime"`

}
