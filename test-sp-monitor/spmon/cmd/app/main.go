package main

import (
	"html/template"
	"os"
	"sp-monitoring/database"
	"sp-monitoring/handlers"

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

	//initialize gin
	r := gin.Default()
	//get database type to be able to use methods related to redis
	rCache := &database.RedisCache{}
	//trying to connect to the reidis
	rCache.RedisConnect(redisHost, redisPort, redisPassword)
	cl := rCache.GetRedisClient()
	//some monitoring if we nned as middleware
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.Request.Header.Get("X-Debug-Key")+" "+param.Request.Header.Get("X-Debug-JSON")+" "+param.Request.Header.Get("X-Debug-Key1"),
	// 	)
	// }))
	// passing the redis cache to the handlers
	hand := handlers.NewHandlers(r, cl)

	// Load HTML template

	r.SetHTMLTemplate(template.Must(template.New("index").Funcs(template.FuncMap{
		"mod": mod,
	}).ParseFiles("templates/index.html")))
	//r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/index.html")).Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }}))
	r.GET("/", hand.GetHandler)
	r.POST("/add", hand.AddHandler)
	r.POST("/update", hand.UpdateHandler)

	r.POST("/delete", hand.DeleteHandler)

	// Run the server
	r.Run(":8080")

}

func mod(i, j int) bool {
	return i%j == 0
}
