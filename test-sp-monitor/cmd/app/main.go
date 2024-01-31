package main

import (
	"fmt"
	"html/template"
	"os"
	"test-sp-monitor/database"
	"test-sp-monitor/handlers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	redisHost := os.Getenv("REDIS_HOST")

	if len(redisHost) == 0 {
		redisHost = "localhost"
	}
	redisPort := os.Getenv("REDIS_PORT")
	if len(redisPort) == 0 {
		redisPort = "6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	if len(redisPassword) == 0 {
		redisPassword = ""
	}

	r := gin.Default()
	rCache := &database.RedisCache{}
	rCache.RedisConnect(redisHost, redisPort, redisPassword)
	cl := rCache.GetRedisClient()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.Request.Header.Get("X-Debug-Key")+" "+param.Request.Header.Get("X-Debug-JSON")+" "+param.Request.Header.Get("X-Debug-Key1"),
		)
	}))

	hand := handlers.NewHandlers(r, cl)

	// Load HTML template
	r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))
	r.GET("/", hand.GetHandler)
	r.POST("/add", hand.AddHandler)

	r.POST("/delete", hand.DeleteHandler)

	// Run the server
	r.Run(":8080")

}
