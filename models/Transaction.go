package models

import "gopkg.in/mgo.v2/bson"

type Transaction struct {
	Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
	WalletId        bson.ObjectId `bson:"walletId,omitempty" json:"walletId"`
	Amount          float32 `json:"amount" form:"amount"`
	TransactionType string`json:"type" form:"type"`
	unixDateTime    int64 `json:"unixDateTime" form:"unixDateTime"`
}
