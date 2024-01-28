package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// PageData represents the data structure for the HTML template
type KVKey struct {
	Secret      string `json:"key" redis:"secret"`
	Keyvault    string `json:"value" redis:"keyvault"`
	Expireddate string `json:"thirdvalue" redis:"expireddate"`
}

type PageData struct {
	PageDataArray []KVKey
}

type Handlers struct {
	r  *gin.Engine
	db *database.RedisCache
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
		PageDataArray: kvkeys,
	}

	// Render the HTML page with the PageData struct
	c.HTML(http.StatusOK, "index.html", gin.H{
		"PageData": pageData,
	})

}

func (h *Handlers) AddHandler(c *gin.Context) {

	// Get key and value from the form
	secret := c.PostForm("secret")
	keyvault := c.PostForm("keyvault")
	//we generate dummy data value
	expireddate, err := displaySecretExpiration(secret, keyvault)
	if err != nil {
		expireddate = ""
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		// Redirect to the main page after adding the album
		c.Redirect(http.StatusSeeOther, "/")
		return


	}

	kvkey := KVKey{
		Secret:      secret,
		Keyvault:    keyvault,
		Expireddate: expireddate,
	}
	// Convert the Album struct to JSON
	albumJSON, err := json.Marshal(kvkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique key for the album
	newkey := fmt.Sprintf("KvKeys:%s", strings.ToLower(secret))

	// // Log the key and JSON data for debugging
	// c.Request.Header.Add("X-Debug-Key", newkey)
	// c.Request.Header.Add("X-Debug-Key1", expireddate.Format("2006-01-02 15:04:05"))
	// c.Request.Header.Add("X-Debug-JSON", string(albumJSON))

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
	secret := c.PostForm("secret")

	// Generate the key for the album
	newkey := fmt.Sprintf("KvKeys:%s", strings.ToLower(secret))

	// Delete the key from the cache


	// Delete the secret from the Redis cache
	err := h.client.Del(newkey).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the home page after deleting the key
	c.Redirect(http.StatusSeeOther, "/")

}

// // just for testing
// func randate() time.Time {
// 	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
// 	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
// 	delta := max - min

// 	sec := rand.Int63n(delta) + min
// 	return time.Unix(sec, 0)
//}
