package main

import (
	"github.com/go-redis/redis"
)

func insertTask(domain string, message string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}

	hSetErr := client.HSet(redis_set_key, domain, message)
	if hSetErr != nil {
		return false
	}

	return true
}

func getTaskByDomain(domain string) string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return ""
	}
	result,_:=client.HGet(redis_set_key,domain).Result()
	return result
}