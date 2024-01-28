package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	RedisConnect()
}

type RedisCache struct {
	Client *redis.Client
}

func (r *RedisCache) RedisConnect(redisHost string, redisPort string, redisPassword string) *RedisCache {
	// Initialize the Redis client
	cl := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// Check if Redis is reachable
	pong, err := connectToDb(cl)

	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	} else {
		fmt.Println("Connected to Redis:", pong)
	}
	return &RedisCache{
		Client: cl,
	}

}

func connectToDb(r *redis.Client) (string, error) {
	var counts int64

	for {
		// Check if Redis is reachable
		pong, err := r.Ping().Result()
		if err != nil {
			log.Println("Redis server is not yet ready")
			counts++
		} else {
			return pong, nil

		}

		if counts > 10 {
			log.Println(err)
			return "", errors.New("could not connect to the redis")
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue

	}
}
