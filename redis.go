package main

import "github.com/garyburd/redigo/redis"

func NewRedis() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	HandleError(err)
	return c
}
