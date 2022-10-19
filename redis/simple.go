// docker run --name redis_test -d -p 6379:6379 redis

package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
	return client
}

func main() {
	client := NewClient()
	if err := client.Set("key", "value", 0).Err(); err != nil {
		log.Println("client.Set failed", err)
		return
	}

	val, err := client.Get("key").Result()
	if err != nil {
		fmt.Printf("client.Get failed: %v\n", err)
		return
	}

	fmt.Printf("val = %v\n", val)
	client.Del("key")

	val2, err := client.Get("key").Result()
	if err == redis.Nil {
		fmt.Println("val does not exist")
	} else if err != nil {
		fmt.Printf("client.Get failed: %v\n", err)
		return
	} else {
		fmt.Printf("key = %v\n", val2)
	}
}
