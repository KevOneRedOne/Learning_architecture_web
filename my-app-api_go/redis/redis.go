package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Init() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "root",
		DB:       0,
	})

	if err := redisClient.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	if pong := redisClient.Ping(ctx); pong.String() != "ping: PONG" {
		log.Println("-------------Error connection redis ----------:", pong)
	}

}
