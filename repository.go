package main

import (
	"github.com/go-redis/redis"
)

func insertTask(key string,field string,value string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}

	hSetErr := client.HSet(key, field, value)
	if hSetErr != nil {
		return false
	}

	return true
}

func getTaskByDomain(key string,field string) string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return ""
	}
	result,_:=client.HGet(key,field).Result()
	return result
}