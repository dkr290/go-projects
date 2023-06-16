package utils

import (
	"encoding/json"
	"fmt"
	"go-acr-clean/pkg/domain"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	ACR          string
	EncCreds     string
	SelectedRepo string
}

func NewConfig(enc string, ACR string, repo string) *Config {
	return &Config{

		ACR:          ACR,
		EncCreds:     enc,
		SelectedRepo: repo,
	}
}

func (c *Config) ListAllImages() []string {

	r, err := http.NewRequest("GET", "https://"+c.ACR+"/v2/_catalog", nil) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Basic "+c.EncCreds)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		//handle error
		log.Fatal(`error: `, err)
	}

	log.Println(res.Status)

	body, error := ioutil.ReadAll(res.Body)

	if error != nil {
		fmt.Println(error)
	}
	// close response body
	res.Body.Close()

	var rp domain.Repos
	err = json.Unmarshal(body, &rp)
	if err != nil {
		panic(err)
	}
	return rp.Repositories
}

func (c *Config) RetarnAllTags() []string {

	//r, err := http.NewRequest("GET", "https://"+ACR+"/v2/"+v+"/tags/list", nil) // URL-encoded payload
	r, err := http.NewRequest("GET", "https://"+c.ACR+"/v2/"+c.SelectedRepo+"/tags/list", nil) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	//r.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Basic "+c.EncCreds)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		//handle error
		log.Fatal(`error: `, err)
	}

	log.Println(res.Status)

	body, error := ioutil.ReadAll(res.Body)

	if error != nil {
		fmt.Println(error)
	}
	// // close response body
	res.Body.Close()
	var rt domain.RegistryTag
	err = json.Unmarshal(body, &rt) //
	if err != nil {
		panic(err)
	}

	return rt.Tags

}
