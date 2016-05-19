package store

import (
	"github.com/garyburd/redigo/redis"
)

type RedisCli struct {
	conn redis.Conn
}

var instanceRedisCli *RedisCli = nil

func Connect() (conn *RedisCli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(RedisCli)
		var err error
		//this is works!!!
		instanceRedisCli.conn, err = redis.Dial("tcp", ":6379")

		if err != nil {
			panic(err)
		}

		if _, err := instanceRedisCli.conn.Do("AUTH", "1"); err != nil {
			instanceRedisCli.conn.Close()
			panic(err)
		}

	}
	return instanceRedisCli
}

func (redisCli *RedisCli) SetValue(key, value string, expiration ...interface{}) error {
	instanceRedisCli = Connect()
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}
	instanceRedisCli = nil
	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	instanceRedisCli = Connect(); //I know that this is strange
	data, err := redisCli.conn.Do("GET", key)
	if err != nil{
		panic(err)
	}
	instanceRedisCli = nil
	return data, err
}
