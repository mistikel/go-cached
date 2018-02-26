package main

import "github.com/garyburd/redigo/redis"

func NewRedis() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}
