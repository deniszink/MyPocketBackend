package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
)

type MongoDB struct {
	session *mgo.Session
	mongodb *mgo.Database
}

var TableUsers string = "users"
var TableWallets string = "wallets"
var mongoInstance *MongoDB

const (
	MongoDBHosts = "ds011271.mlab.com:11271"
	AuthDatabase = "heroku_96lfkqsw"
	AuthUserName = "denisz"
	AuthPassword = "mypocket"

	/*MongoDBHosts = "127.0.0.1:27017"
	AuthDatabase = "mypocket"
	AuthUserName = ""
	AuthPassword = ""*/
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

		mongoInstance.mongodb = mongoInstance.session.DB(AuthDatabase)

		mongoInstance.mongodb.C(TableUsers)
		mongoInstance.mongodb.C(TableWallets)

		if err != nil {
			panic(err)
		}

	}

	return mongoInstance
}

func (this *MongoDB) CreateTable(name string) (c *mgo.Collection) {
	return this.mongodb.C(name)
}

func (this *MongoDB) WriteDataTo(tableName string, data interface{}) error {
	fmt.Println(data)
	table := this.mongodb.C(tableName)
	return table.Insert(data)
}

func (this *MongoDB) FindAll(tableName string, selector bson.M, source interface{}) error {
	table := this.mongodb.C(tableName)
	return table.Find(selector).All(source)
}

func (this *MongoDB) FindOne(tableName string, selector bson.M, source interface{}) error {
	table := this.mongodb.C(tableName)
	return table.Find(selector).One(source)
}
func (this *MongoDB) IsExists(tableName string, selector bson.M)  (bool,error) {
	table := this.mongodb.C(tableName)
	count, err := table.Find(selector).Count()
	fmt.Print(count, err)
	return count > 0, err
}

func (this *MongoDB) GetOne(tableName string, selector bson.M, source interface{}) error {
	table := this.mongodb.C(tableName)
	return table.Find(selector).One(source)
}

func (this *MongoDB) Update(tableName string, model interface{}, source interface{}) {
	//todo implement change email for user
}


