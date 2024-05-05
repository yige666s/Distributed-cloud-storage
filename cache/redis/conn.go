package redis

import (
	"fmt"
	"time"

	// "github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

var (
	pool      *redis.Pool
	redisHost = "23.95.130.120:6380" // 注意修改IP和端口占用情况
	redisPass = "password"
)

// 创建Redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50, // 最大连接数
		MaxActive:   30, // 同时连接数
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) { //创建连接的方法
			// 1.打开连接
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			// 2. 访问认证
			if _, err = c.Do("AUTH", redisPass); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, lastUsed time.Time) error { // 定期检测redis可用性，出错则断开
			if time.Since(lastUsed) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
