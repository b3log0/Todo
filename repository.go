package main

import (
	"github.com/go-redis/redis"
)
func setCurrentDomain(domain string){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err == nil {
		client.Set(CURRENT_KEY,domain,0)
	}
}

func getCurrentDomain() string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return ""
	}
	result,_ := client.Get(CURRENT_KEY).Result()
	return result 
}

func getDomain(index int64) string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return ""
	}
	result,_:=client.LIndex(REDIS_KEY,index).Result()
	return result
}

func insertDomain(domain string) bool{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}
	result,_ := client.RPush(REDIS_KEY,domain).Result()
	if result > 0 {
		return true
	}else{
		return false
	}
}

func getDomains() []string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err == nil {
		result,_ := client.LRange(REDIS_KEY,0,-1).Result()
		return result
	}else{
		return nil
	}
}

func delDomain(domain string) bool{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}
	result,_ := client.LRem(REDIS_KEY,0,domain).Result()
	if result > 0 {
		return true
	}else{
		return false
	}
}

func setTask(domain string,task string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}
	result,_ := client.RPush(domain,task).Result()
	if result > 0 {
		return true
	}else{
		return false
	}
}

func getTask(domain string,index int64) string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return ""
	}
	result,_:=client.LIndex(domain,index).Result()
	return result
}

func getTasks(domain string) []string{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err == nil {
		result,_:=client.LRange(domain,0,-1).Result()
		return result
	}else{
		return nil
	}
}

func delTask(domain string,value string) bool{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return false
	}
	result,_ := client.LRem(domain,0,value).Result()
	if result > 0 {
		return true
	}else{
		return false
	}
}