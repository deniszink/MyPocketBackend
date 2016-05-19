package models

import "gopkg.in/mgo.v2/bson"

type Wallet struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	WalletName string `json:"walletName" form:"walletName"`
	Currency   string `json:"currency" form:"currency"`
	Balance    float64 `json:"balance" form:"balance"`
	UserID     bson.ObjectId `bson:"userId,omitempty" json:"userId" form:"userId"`
}
