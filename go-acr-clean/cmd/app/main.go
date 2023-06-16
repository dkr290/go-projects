package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"go-acr-clean/pkg/utils"
	"log"
	"os"
	"sort"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")

	}

	enc := encoderCreds()
	ACR := os.Getenv("ACR")

	FlagSetList := flag.NewFlagSet("FlagSetList", flag.ContinueOnError)

	ListAll := FlagSetList.Bool("list-all-repos", false, "List all images for the ACR")
	RepoTags := FlagSetList.String("list-repo-tags", "", "Pass the repository from the list with --list-all")

	FlagSetDeleteTagsFromTo := flag.NewFlagSet("FlagSetDeleteTagsFromTo", flag.ContinueOnError)
	deletetags := FlagSetDeleteTagsFromTo.Bool("delete-tags", false, "")
	repo := FlagSetDeleteTagsFromTo.String("repo", "", "")
	starttag := FlagSetDeleteTagsFromTo.String("start-tag", "", "")
	endtag := FlagSetDeleteTagsFromTo.String("end-tag", "", "")

	pflag.CommandLine.AddGoFlagSet(FlagSetList)
	pflag.CommandLine.AddGoFlagSet(FlagSetDeleteTagsFromTo)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	switch {
	case *ListAll:
		conf := utils.NewConfig(enc, ACR, *RepoTags)
		repos := conf.ListAllImages()
		fmt.Println("===================")
		TablePrint(repos)
	case *RepoTags != "":
		conf := utils.NewConfig(enc, ACR, *RepoTags)
		tags := conf.CalculateTags()
		sortTagsByTimestamp(tags)
	case *deletetags:
		if *repo != "" && *starttag != "" && *endtag != "" {
			conf := utils.NewConfig(enc, ACR, *repo)
			log.Println(conf.SelectedRepo)
			alltags := conf.CalculateTags()
			conf.ScliceTagsStartEnd(alltags, *starttag, *endtag)

		} else {
			fmt.Println("Please use appropriate options")
			fmt.Println("go-acr-clean or go-acr-clean.exe --delete-tags --repo <repository name> --start-tag <the oldest tag to delete> --end-tag <the newwest tag to delete>")
			return
		}
	default:
		fmt.Println("Please use appropriate options")
		fmt.Println("go-acr-clean or go-acr-clean.exe --list-all-repos   to list all repositories")
		fmt.Println("go-acr-clean or go-acr-clean.exe --list-repo-tags   <repository name> list all tags fopr the repository")
		fmt.Println("go-acr-clean or go-acr-clean.exe --delete-tags --repo <repository name> --start-tag <the oldest tag to delete> --end-tag <the newwest tag to delete>")
	}

}

func encoderCreds() string {

	USERNAME := os.Getenv("ACR_USERNAME")
	PASSWORD := os.Getenv("ACR_PASSWORD")

	data := []byte(USERNAME + ":" + PASSWORD)
	enc := base64.StdEncoding.EncodeToString(data)
	return enc
}

func TablePrint(repos []string) {

	for i, k := range repos {
		fmt.Println(i, k)

	}
}

func sortTagsByTimestamp(m map[string]time.Time) {

	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]].Before(m[keys[j]])
	})

	for key, _ := range keys {
		fmt.Printf("%v  \t\t%v\n", keys[key], m[keys[key]])

	}

}
