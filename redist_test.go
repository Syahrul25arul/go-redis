package goredis

import (
	"log"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestConnectRedis(t *testing.T) {
	conn, err := redis.Dial("tcp", ":6379")
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}

	_, err = conn.Do("SET", "test5", "test toh")
	if err != nil {
		log.Panic(err.Error())
	}
}
