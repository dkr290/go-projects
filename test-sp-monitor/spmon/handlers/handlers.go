package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
	"sync"

	"sp-monitoring/helpers"
	"sp-monitoring/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Handlers struct {
	r        *gin.Engine
	client   *redis.Client
	userName string
	password string
}

// defining the New function as factory pattern
func NewHandlers(r *gin.Engine, rc *redis.Client, user, pass string) *Handlers {
	return &Handlers{
		r:        r,
		client:   rc,
		userName: user,
		password: pass,
	}
}

// Handler function for the login page
func (h *Handlers) LoginPageHandler(c *gin.Context) {
	var tmpl = template.Must(template.ParseFiles("templates/login.html"))

	// Render the login page template
	err := tmpl.ExecuteTemplate(c.Writer, "login.html", nil)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}

// Define routes
func (h *Handlers) GetHandler(c *gin.Context) {
	if !isAuthenticated(c) {
		// Redirect to the login page if not authenticated
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Get the current page number from the query parameters
	pageStr := c.Request.URL.Query().Get("page")
	// Get the search query from the query parameters
	searchQuery := c.Request.URL.Query().Get("search")

	// Fetch all keys from the Redis cache
	keys, err := h.client.Keys("KvKeys:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice to store secrets
	var kvkeys []models.DynamicData

	// Fetch secrets for each key
	for _, key := range keys {
		kvkeysJSON, err := h.client.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var kvkey models.KVKey
		err = json.Unmarshal([]byte(kvkeysJSON), &kvkey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		duration, err := helpers.CheckTime(kvkey)
		if err != nil {
			kvkeys = append(kvkeys, models.DynamicData{WarningMessage: 3, KVKey: kvkey})
			continue
		}
		if duration.Hours() < 30*24 {
			kvkeys = append(kvkeys, models.DynamicData{WarningMessage: 1, KVKey: kvkey})

		} else if duration.Hours() < 60*24 {
			kvkeys = append(kvkeys, models.DynamicData{WarningMessage: 2, KVKey: kvkey})

		} else {
			kvkeys = append(kvkeys, models.DynamicData{WarningMessage: 3, KVKey: kvkey})
		}

	}
	// Sort the Pagerdata by Expireddate (as string)
	sort.Slice(kvkeys, func(i, j int) bool {
		return kvkeys[i].Expireddate < kvkeys[j].Expireddate
	})

	totalItems := len(kvkeys)
	itemsPerPage := 10
	pagination, CurrentData := FingCurrentPage(pageStr, totalItems, itemsPerPage, kvkeys)

	var pageData models.PageData
	var searchedOutput []models.DynamicData

	// Get the search query from the query parameters
	if searchQuery != "" {
		searchedOutput = SearchPage(searchQuery, kvkeys)
		// Create a PageData struct
		pageData = models.PageData{
			PageDataArray: searchedOutput,
			Pagination:    pagination,
			SearchQuery:   searchQuery,
		}

	} else {
		// Create a PageData struct
		pageData = models.PageData{
			PageDataArray: CurrentData,
			Pagination:    pagination,
			SearchQuery:   searchQuery,
		}

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
	metadata := c.PostForm("metadata")
	//we generate dummy data value
	expireddate, err := helpers.DisplaySecretExpiration(secret, keyvault)
	if err != nil {
		expireddate = ""
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		// Redirect to the main page after adding the album
		c.Redirect(http.StatusSeeOther, "/")
		return

	}

	kvkey := models.KVKey{
		Secret:      secret,
		Keyvault:    keyvault,
		Expireddate: expireddate,
		Metadata:    metadata,
	}
	// Convert the KVKey  struct to JSON
	kvkeysJSON, err := json.Marshal(kvkey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kvName, err := helpers.ExtractKVName(keyvault)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique key for the album
	newkey := fmt.Sprintf("KvKeys:%s:%s", strings.ToLower(secret), strings.ToLower(kvName))

	// // Log the key and JSON data for debugging
	// c.Request.Header.Add("X-Debug-Key", newkey)
	// c.Request.Header.Add("X-Debug-Key1", expireddate.Format("2006-01-02 15:04:05"))
	// c.Request.Header.Add("X-Debug-JSON", string(albumJSON))

	// Add the album to the Redis cache
	err = h.client.Set(newkey, kvkeysJSON, 0).Err()
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
	keyvault := c.PostForm("keyvault")

	// Generate the key for the album
	kvName, err := helpers.ExtractKVName(keyvault)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newkey := fmt.Sprintf("KvKeys:%s:%s", strings.ToLower(secret), strings.ToLower(kvName))

	// Delete the key from the cache

	// Delete the album from the Redis cache
	err = h.client.Del(newkey).Err()
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

func (h *Handlers) UpdateHandler(c *gin.Context) {
	// Fetch all keys from the Redis cache
	errorChan := make(chan CustomError, 1)
	keys, err := h.client.Keys("KvKeys:*").Result()
	var wg sync.WaitGroup

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to fetch all keys from redis",
		})
		return
	}
	// Start a goroutine to close the error channel once all goroutines are done
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	for _, key := range keys {
		wg.Add(1)
		go h.checkKvKeysRedis(key, errorChan, &wg)
	}

	// Collect and handle errors from the error channel
	for customErr := range errorChan {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   customErr.Err.Error(),
			"message": customErr.Message,
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

// Handler function for the login form submission
func (h *Handlers) LoginHandler(c *gin.Context) {

	// Get the username and password from the form
	submittedUsername := c.PostForm("username")
	submittedPassword := c.PostForm("password")

	// Check if the username and password are correct
	if submittedUsername == h.userName && submittedPassword == h.password {
		// Set a session cookie to indicate that the user is authenticated
		c.SetCookie("authenticated", "true", 3600, "/", "", false, true)

		// Redirect to the home page
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	var tmpl = template.Must(template.ParseFiles("templates/login.html"))
	// Render the login page with an error message
	err := tmpl.ExecuteTemplate(c.Writer, "login.html", "Invalid username or password")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}

// Helper function to check if the user is authenticated
func isAuthenticated(c *gin.Context) bool {
	// Check if the "authenticated" cookie is present
	_, err := c.Cookie("authenticated")
	return err == nil
}

type CustomError struct {
	Err     error
	Message string
}

// goroutines to check and compare from redis when update is clicked for changes in keyvault expiration
func (h *Handlers) checkKvKeysRedis(key string, errorChan chan CustomError, wg *sync.WaitGroup) {
	defer wg.Done()
	kvkeysJSON, err := h.client.Get(key).Result()
	// if err == nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{

	// 		"message":  key,
	// 		"message1": kvkeysJSON,
	// 	})
	// }
	if err != nil {
		errorChan <- CustomError{Err: err, Message: "failed to fetch key in json"}
		return
	}
	// unmarshal
	var kvkey models.KVKey
	err = json.Unmarshal([]byte(kvkeysJSON), &kvkey)
	if err != nil {
		errorChan <- CustomError{Err: err, Message: "error unmarshaling items"}
		return
	}
	//helper function to check for secrets expiration from the keyvault
	expireddate, err := helpers.DisplaySecretExpiration(kvkey.Secret, kvkey.Keyvault)
	if err != nil {
		errorChan <- CustomError{Err: err, Message: "error fetch secret expiration"}
		return
	}
	fromRedis, fromKV, err := helpers.ConvertToTime(kvkey, expireddate)
	if err != nil {
		errorChan <- CustomError{Err: err, Message: "Error convert to time"}
		return
	}

	if fromRedis != fromKV {
		kvkey.Expireddate = expireddate
		kvName, err := helpers.ExtractKVName(kvkey.Keyvault)
		if err != nil {
			errorChan <- CustomError{Err: err, Message: "error extracting keyvault name"}
			return
		}
		newkey := fmt.Sprintf("KvKeys:%s:%s", strings.ToLower(kvkey.Secret), strings.ToLower(kvName))

		kvKeysJSON, err := json.Marshal(kvkey)
		if err != nil {
			errorChan <- CustomError{Err: err, Message: "error marshaling the new values"}
			return
		}
		err = h.client.Set(newkey, kvKeysJSON, 0).Err()
		if err != nil {
			errorChan <- CustomError{Err: err, Message: "error saving the new value to redis"}
			return
		}

	}
}
