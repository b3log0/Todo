/*
 * Redis 操作
 * 未必启用，将来可能在查询或者读取逻辑变得复杂时改为redis操作
 * 即在读命令时通过redis，写命令时写入文件
 */
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := c.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get from Redis:>", v)
	defer c.Close()
}
