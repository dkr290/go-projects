package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	// Initialize the Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change this to your Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Check if Redis is reachable
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
	} else {
		fmt.Println("Connected to Redis:", pong)
	}
}

// PageData represents the data structure for the HTML template
type PageData struct {
	KeyValues map[string]string
}

type Handlers struct {
	r *gin.Engine
}

func NewHandlers(r *gin.Engine) *Handlers {
	return &Handlers{
		r: r,
	}
}

// Define routes
func (h *Handlers) GetHandler(c *gin.Context) {

	// Retrieve all key-value pairs from the cache
	keyValues, err := client.HGetAll("myCache").Result()
	if err != nil && err != redis.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render HTML template with key-value pairs
	pageData := PageData{
		KeyValues: keyValues,
	}
	c.HTML(http.StatusOK, "index.html", pageData)

}

func (h *Handlers) AddHandler(c *gin.Context) {

	// Get key and value from the form
	key := c.PostForm("key")
	value := c.PostForm("value")

	// Add the key-value pair to the cache
	err := client.HSet("myCache", key, value).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the home page after adding the key-value pair
	c.Redirect(http.StatusSeeOther, "/")

}

func (h *Handlers) DeleteHandler(c *gin.Context) {

	// Get key to delete from the form
	key := c.PostForm("key")

	// Delete the key from the cache
	err := client.HDel("myCache", key).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the home page after deleting the key
	c.Redirect(http.StatusSeeOther, "/")

}
