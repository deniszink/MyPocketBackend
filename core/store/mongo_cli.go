package store

import (
	"github.com/go-mgo/mgo"
	"github.com/go-mgo/mgo/bson"
)

type MongoDB struct {
	session *mgo.Session
	mongodb *mgo.Database
}

var TableUsers string = "users"

var mongoInstance *MongoDB

func ConnectMongo() (mongo *MongoDB){
	if mongoInstance == nil{
		mongoInstance = new(MongoDB)

		var err error
		mongoInstance.session,err = mgo.Dial("")
		defer  mongoInstance.session.Close()

		if err != nil{
			panic(err)
		}

		mongoInstance.session.SetMode(mgo.Monotonic,true)

		mongoInstance.mongodb = mongoInstance.session.DB("mypocket_db")

		mongoInstance.mongodb.C(TableUsers)



	}

	return mongoInstance
}

func(this *MongoDB) CreateTable(name string)(c *mgo.Collection){
	return this.mongodb.C(name)
}

func(this *MongoDB) WriteDataTo(tableName string, data *struct{}){

	table := this.mongodb.C(tableName)
	err := table.Insert(&data)
	if err != nil{
		panic(err)
	}
}

func(this *MongoDB) FindAll(tableName string,selector bson.M, source interface{}){
	table := this.mongodb.C(tableName)
	table.Find(selector).All(&source)
}

func(this *MongoDB) FindOne(tableName string, selector bson.M, source interface{}){
	table := this.mongodb.C(tableName)
	table.Find(selector).All(&source)
}

func(this *MongoDB) Update(tableName string, model interface{}, source interface{}){

}


