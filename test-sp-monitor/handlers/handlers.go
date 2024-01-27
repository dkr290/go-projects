package handlers

import (
	"net/http"
	"test-sp-monitor/database"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// PageData represents the data structure for the HTML template
type PageData struct {
	KeyValues map[string]string
}

type Handlers struct {
	r  *gin.Engine
	db *database.RedisCache
}

func NewHandlers(r *gin.Engine, redis *database.RedisCache) *Handlers {
	return &Handlers{
		r:  r,
		db: redis,
	}
}

// Define routes
func (h *Handlers) GetHandler(c *gin.Context) {

	// Retrieve all key-value pairs from the cache
	keyValues, err := h.db.Client.HGetAll("myCache").Result()
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
	err := h.db.Client.HSet("myCache", key, value).Err()
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
	err := h.db.Client.HDel("myCache", key).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the home page after deleting the key
	c.Redirect(http.StatusSeeOther, "/")

}
