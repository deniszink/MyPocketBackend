package store

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	prodServer = "lab.redistogo.com:9951"
	prodPassword = "4e82903dfe08366aac967296747c44c8"
	testServer = ":6379"
	testPassword = "1"
)

type redisRequestStore struct {
	redis.Conn
}
type redisTokenGenerator struct {
	tokenKey string
	p *redis.Pool
}

var instanceRedisCli *redisRequestStore = nil

func Connect()  *redisRequestStore {
	 if instanceRedisCli == nil{
		 pool := newRedisPool(prodServer,prodPassword)
		 //pool := newRedisPool(testServer,testPassword)
		 defer pool.Close()
		 instanceRedisCli = NewRedisRequestStore(pool.Get())
		 //tokenGenerator := NewRedisTokenGenerator(conn, "_token")

	 }
	return instanceRedisCli
}

func newRedisPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewRedisRequestStore(conn redis.Conn) *redisRequestStore {
	return &redisRequestStore{conn}
}

func (r redisRequestStore) Get(key string) (string, error) {
	return redis.String(r.Do("GET", key))
}

func (r redisRequestStore) Put(key, value string, expiration ...interface{}) (interface{},error){
	reply,err := r.Do("SET", key, value)

	if err == nil && expiration != nil {
		r.Conn.Do("EXPIRE", key, expiration[0])
	}
	return reply,err
}

func (r redisRequestStore) Delete(key string) error {
	_, err := r.Do("DEL", key)
	return err
}

/*
func NewRedisTokenGenerator(pool *redis.Pool, tokenKey string) *redisTokenGenerator {
	return &redisTokenGenerator{
		tokenKey: tokenKey,
		p: pool,
	}
}

func (r redisTokenGenerator) Next() (int, error) {
	c := r.p.Get()
	defer c.Close()
	token, err := redis.Int(c.Do("INCR", r.tokenKey))
	if err != nil {
		return 0, err
	}
	return token, nil
}*/
