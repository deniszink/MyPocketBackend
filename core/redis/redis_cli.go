package redis

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

		instanceRedisCli.conn, err = redis.Dial("tcp", "redis://redistogo:4e82903dfe08366aac967296747c44c8@lab.redistogo.com:9951")

		if err != nil {
			panic(err)
		}

		if _, err := instanceRedisCli.conn.Do("AUTH", "mypass"); err != nil {
			instanceRedisCli.conn.Close()
			panic(err)
		}
	}

	return instanceRedisCli
}

func (redisCli *RedisCli) SetValue(key, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}
	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}
