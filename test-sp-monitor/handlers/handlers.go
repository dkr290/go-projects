package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// var (
// 	client        *redis.Client
// 	redisHost     = "localhost"
// 	redisPort     = "6379"
// 	redisPassword = ""
// )

// func init() {
// 	// Initialize the Redis client
// 	client = redis.NewClient(&redis.Options{
// 		Addr:     redisHost + ":" + redisPort,
// 		Password: redisPassword,
// 		DB:       0,
// 	})

// 	// Check if Redis is reachable
// 	pong, err := connectToDb(client)

// 	if err != nil {
// 		fmt.Println("Error connecting to Redis:", err)
// 	} else {
// 		fmt.Println("Connected to Redis:", pong)
// 	}
// }

// PageData represents the data structure for the HTML template
type KVKey struct {
	Key        string    `json:"key" redis:"key"`
	Value      string    `json:"value" redis:"value"`
	ThirdValue time.Time `json:"thirdvalue" redis:"thirdvalue"`
}

type PageData struct {
	Albums []KVKey
}

type Handlers struct {
	r      *gin.Engine
	client *redis.Client
}

func NewHandlers(r *gin.Engine, rc *redis.Client) *Handlers {
	return &Handlers{
		r:      r,
		client: rc,
	}
}

// Define routes
func (h *Handlers) GetHandler(c *gin.Context) {

	// Fetch all keys from the Redis cache
	keys, err := h.client.Keys("KvKeys:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice to store albums
	var kvkeys []KVKey

	// Fetch albums for each key
	for _, key := range keys {
		kvkeysJSON, err := h.client.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var kvkey KVKey
		err = json.Unmarshal([]byte(kvkeysJSON), &kvkey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		kvkeys = append(kvkeys, kvkey)
	}

	// Create a PageData struct
	pageData := PageData{
		Albums: kvkeys,
	}

	// Render the HTML page with the PageData struct
	c.HTML(http.StatusOK, "index.html", gin.H{
		"PageData": pageData,
	})

}

func (h *Handlers) AddHandler(c *gin.Context) {

	// Get key and value from the form
	key := c.PostForm("key")
	value := c.PostForm("value")
	//we generate dummy data value
	thirdValue := randate()

	kvkey := KVKey{
		Key:        key,
		Value:      value,
		ThirdValue: thirdValue,
	}
	// Convert the Album struct to JSON
	albumJSON, err := json.Marshal(kvkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique key for the album
	newkey := fmt.Sprintf("KvKeys:%s", strings.ToLower(key))

	// Log the key and JSON data for debugging
	c.Request.Header.Add("X-Debug-Key", newkey)
	c.Request.Header.Add("X-Debug-Key1", thirdValue.Format("2006-01-02 15:04:05"))
	c.Request.Header.Add("X-Debug-JSON", string(albumJSON))

	// Add the album to the Redis cache
	err = h.client.Set(newkey, albumJSON, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the main page after adding the album
	c.Redirect(http.StatusSeeOther, "/")

}

func (h *Handlers) DeleteHandler(c *gin.Context) {

	// Get key to delete from the form
	key := c.PostForm("key")

	// Generate the key for the album
	newkey := fmt.Sprintf("KvKeys:%s", strings.ToLower(key))

	// Delete the key from the cache

	// Delete the album from the Redis cache
	err := h.client.Del(newkey).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the home page after deleting the key
	c.Redirect(http.StatusSeeOther, "/")

}

// just for testing
func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// func connectToDb(r *redis.Client) (string, error) {
// 	var counts int64

// 	for {
// 		// Check if Redis is reachable
// 		pong, err := r.Ping().Result()
// 		if err != nil {
// 			log.Println("Redis server is not yet ready")
// 			counts++
// 		} else {
// 			return pong, nil

// 		}

// 		if counts > 10 {
// 			log.Println(err)
// 			return "", errors.New("could not connect to the redis")
// 		}

// 		log.Println("Backing off for two seconds...")
// 		time.Sleep(2 * time.Second)
// 		continue

// 	}
// }
