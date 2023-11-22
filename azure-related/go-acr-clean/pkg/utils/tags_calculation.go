package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func (c *Config) CalculateTags() map[string]time.Time {
	alltags := c.RetarnAllTags()

	type Tag struct {
		Name           string
		Digest         string
		CreatedTime    string
		LastUpdateTime string
	}

	type Tags struct {
		Registry  string
		ImageName string
		Tag       Tag
	}
	curr_tags := make(map[string]time.Time)

	for _, t := range alltags {

		r, err := http.NewRequest("GET", "https://"+c.ACR+"/acr/v1/"+c.SelectedRepo+"/_tags/"+t, nil) // URL-encoded payload
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
		r.Header.Add("Authorization", "Basic "+c.EncCreds)
		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			//handle error
			log.Fatal(`error: `, err)
		}

		body, error := ioutil.ReadAll(res.Body)

		if error != nil {
			fmt.Println(error)
		}
		// close response body
		res.Body.Close()
		var allTags Tags
		err = json.Unmarshal(body, &allTags)
		if err != nil {
			panic(err)
		}
		//fmt.Println(allTags.Tag.LastUpdateTime)

		re, err := regexp.Compile(`^\d*.*T\d\d:\d\d:\d\d`)
		if err != nil {
			log.Fatal(err)
		}
		matched := re.Find([]byte(allTags.Tag.LastUpdateTime))
		if err != nil {
			log.Fatalln(err)
		}
		all_dates, err := time.Parse("2006-01-02T15:04:05", string(matched))
		if err != nil {
			log.Fatal(err)
		}

		curr_tags[allTags.Tag.Name] = all_dates

	}

	return curr_tags

}
