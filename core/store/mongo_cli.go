package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"backend/models"
	"fmt"
	"time"
)

type MongoDB struct {
	session *mgo.Session
	mongodb *mgo.Database
}

var TableUsers string = "users"

var mongoInstance *MongoDB

const (
	MongoDBHosts = "mlab.com:11271"
	AuthDatabase = "heroku_96lfkqsw"
	AuthUserName = "heroku_96lfkqsw"
	AuthPassword = ""
)

func ConnectMongo() (mongo *MongoDB) {
	if mongoInstance == nil {
		mongoInstance = new(MongoDB)

		var err error
		dialinfo := &mgo.DialInfo{
			Addrs: []string{MongoDBHosts},
			Timeout: 60 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}

		mongoInstance.session, err = mgo.DialWithInfo(dialinfo)

		//defer  mongoInstance.session.Close()

		if err != nil {
			panic(err)
		}

		mongoInstance.session.SetMode(mgo.Monotonic, true)

		mongoInstance.mongodb = mongoInstance.session.DB("mypocket_db")

		mongoInstance.mongodb.C(TableUsers)

		if err != nil {
			panic(err)
		}

	}

	return mongoInstance
}

func (this *MongoDB) CreateTable(name string) (c *mgo.Collection) {
	return this.mongodb.C(name)
}

func (this *MongoDB) WriteDataTo(tableName string, data interface{}) {
	fmt.Println(data)
	table := this.mongodb.C(tableName)
	err := table.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (this *MongoDB) FindAll(tableName string, selector bson.M, source interface{}) error {
	table := this.mongodb.C(tableName)
	return table.Find(selector).All(&source)
}

func (this *MongoDB) FindOne(tableName string, selector bson.M, source interface{}) error {
	user := &models.User{}
	table := this.mongodb.C(tableName)
	return table.Find(selector).One(user)
}

func (this *MongoDB) Update(tableName string, model interface{}, source interface{}) {
	//todo implement change email for user
}


