package main

import (
	"fmt"
	"html/template"
	"test-sp-monitor/database"
	"test-sp-monitor/handlers"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	redisCache := &database.RedisCache{}
	redisCache.RedisConnect("localhost", "6379", "")
	redisClient := redisCache.GetRedisClient()

	hand := handlers.NewHandlers(r, redisClient)
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
	// Load HTML template
	r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")))
	r.GET("/", hand.GetHandler)
	r.POST("/add", hand.AddHandler)
	r.POST("/delete", hand.DeleteHandler)

	// Run the server
	r.Run(":8080")

}
