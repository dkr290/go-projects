package teamsalerts

import (
	"encoding/json"
	"fmt"
	"sp-monitoring/helpers"
	"sp-monitoring/models"
)

func TeamsTask() {

	// Fetch all keys from the Redis cache
	keys, err := client.Keys("KvKeys:*").Result()

	if err != nil {
		fmt.Println("Error: failed to fetch all keys from redis" + err.Error())

	}
	for _, key := range keys {
		kvkeysJSON, err := client.Get(key).Result()

		if err != nil {
			fmt.Println("Error: failed to fetch key in json" + err.Error())

		}

		var kvkey models.KVKey
		err = json.Unmarshal([]byte(kvkeysJSON), &kvkey)
		if err != nil {
			fmt.Printf("error %v", err.Error())

		}
		duration, err := helpers.CheckTime(kvkey)
		if err != nil {
			fmt.Printf("error %v", err.Error())
		}
		if duration.Hours() < 30*24 {
			err := SendAlert("Alert!! the service principal: " + kvkey.Secret + " metadata: " +
				kvkey.Metadata + " in keyvault " + kvkey.Keyvault + " will expire in < 30 days")
			if err != nil {
				fmt.Printf("error %v" + err.Error())
			}

		}

	}

}
