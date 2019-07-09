package db

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// Pool ..
var pool *redis.Pool

// InitPoll ..
func initPoll() {
	maxIdle := viper.GetInt("Redis.maxIdle")
	if maxIdle <= 10 {
		maxIdle = 30
	}
	Pool := NewRedisPool(viper.GetString("Redis.url"), maxIdle, time.Second*240)
	_, err := Pool.Dial()
	if err != nil {
		log.Fatal("redis error: ", err)
	}
	// 清空redis库
	if viper.GetBool("Redis.clearUp") {
		conn := Pool.Get()
		conn.Do("flushdb")
		conn.Close()
	}
	pool = Pool
}

// NewRedisPool returns a new Redis connection pool.
func NewRedisPool(redisURL string, maxIdle int, idleTimeout time.Duration) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL,
				redis.DialReadTimeout(time.Second*3),
				redis.DialWriteTimeout(time.Second*3),
			)
			if err != nil {
				log.Fatal("redis connection error: ", err)
			}
			c.Do("SELECT", 1)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Now().Sub(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

// GetRedis .
func GetRedis() redis.Conn {
	return pool.Get()
}

// // DoRedis ..
// func DoRedis(key, value string, args ...interface{}) {
// 	conn := GetRedis()
// 	defer conn.Close()
// 	conn.Do(key, value, args...)
// }
