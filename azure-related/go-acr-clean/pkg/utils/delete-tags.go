package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

func (c *Config) DeleteDockerImageTag(t string) error {

	r, err := http.NewRequest("DELETE", "https://"+c.ACR+"/acr/v1/"+c.SelectedRepo+"/_tags/"+t, nil) // URL-encoded payload
	if err != nil {
		return err
	}
	r.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	r.Header.Add("Authorization", "Basic "+c.EncCreds)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	log.Println(res.Status)

	_, error := ioutil.ReadAll(res.Body)

	if error != nil {
		return err
	}
	// close response body
	res.Body.Close()

	return nil

}

func (c *Config) ScliceTagsStartEnd(m map[string]time.Time, start string, end string) {

	keys := make([]string, 0, len(m))
	var timeStampStart, timeStampEnd time.Time

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]].Before(m[keys[j]])
	})

	for key, _ := range keys {
		if keys[key] == start {
			timeStampStart = m[keys[key]]

		} else if keys[key] == end {
			timeStampEnd = m[keys[key]]

		}

	}

	for key, _ := range keys {
		if inTimeSpan(timeStampStart, timeStampEnd, m[keys[key]]) {
			err := c.DeleteDockerImageTag(keys[key])
			if err != nil {
				log.Fatalln("Delingt tag fails ", err)
			}
			log.Println("deleted", keys[key], "with timestamp", m[keys[key]])
		}
	}

}

func inTimeSpan(start, end, check time.Time) bool {

	return check.After(start) && check.Before(end)
}
