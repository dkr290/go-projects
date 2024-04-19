package teamsalerts

import "github.com/go-redis/redis"

var webhookURL string
var client *redis.Client

func Get_envs(w string, c *redis.Client) {

	webhookURL = w
	client = c

}
