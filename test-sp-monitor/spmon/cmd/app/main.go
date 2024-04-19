package main

import (
	"html/template"
	"log"
	"os"
	"sp-monitoring/database"
	"sp-monitoring/handlers"
	"sp-monitoring/teamsalerts"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
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

	webhookURL := os.Getenv("TEAMS_WEBHOOK_URL") // Get the webhook URL from environment variable

	if webhookURL == "" {
		log.Fatalln("TEAMS_WEBHOOK_URL environment variable is not set")
	}

	//initialize gin
	r := gin.Default()
	//get database type to be able to use methods related to redis
	rCache := &database.RedisCache{}
	//trying to connect to the reidis
	rCache.RedisConnect(redisHost, redisPort, redisPassword)
	cl := rCache.GetRedisClient()
	r.Use(gin.Recovery())

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

	teamsalerts.Get_envs(webhookURL, cl)

	go func() {
		s := gocron.NewScheduler()
		s.Every(1).Day().Do(teamsalerts.TeamsTask)
		//s.Every(1).Minute().Do(teamsalerts.TeamsTask)
		<-s.Start()
	}()

	r.Use(func(c *gin.Context) {
		c.Next()

		// Check if an error occurred
		err := c.Errors.Last()
		if err != nil {
			// Log the error to the console
			log.Println("Error:", err.Error())
		}
	})

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
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}

}

func mod(i, j int) bool {
	return i%j == 0
}
