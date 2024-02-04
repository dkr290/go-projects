package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

// very interesting method so with interfaces we keep methods and change the type if we need another database type
type Cache interface {
	RedisConnect(string, string, string)
	GetRedisClient()
}

// this is just a type containing the redis client
type RedisCache struct {
	client *redis.Client
}

// RedisConnect(redisHost string, redisPort string, redisPassword string) - initializes the redis client
func (r *RedisCache) RedisConnect(redisHost string, redisPort string, redisPassword string) {
	// Initialize the Redis client
	r.client = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// Check if Redis is reachable
	pong, err := connectToDb(r.client)

	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	} else {
		fmt.Println("Connected to Redis:", pong)
	}

}

// GetRedisClient() return the redis client usually in main to be used everywhere
func (r *RedisCache) GetRedisClient() *redis.Client {
	return r.client
}

// connectToDb - helper function to loop and try to connect to the redis then after a while give up
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
