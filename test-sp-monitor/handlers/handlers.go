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

func NewHandlers(r *gin.Engine, redis *redis.Client) *Handlers {
	return &Handlers{
		r:      r,
		client: redis,
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
	//key := c.PostForm("key")
	key := "test1"
	//value := c.PostForm("value")
	value := "test2"
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
	newkey := fmt.Sprintf("KvKeys:%s:%s", strings.ToLower(key), strings.ToLower(value))

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
	c.Redirect(http.StatusPermanentRedirect, "/")

}

func (h *Handlers) DeleteHandler(c *gin.Context) {

	// Get key to delete from the form
	key := c.PostForm("key")

	// Delete the key from the cache
	err := h.client.HDel("myCache", key).Err()
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
