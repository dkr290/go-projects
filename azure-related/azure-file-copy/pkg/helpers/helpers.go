package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvs(env string) {

	if err := godotenv.Load(env); err != nil {
		log.Println("No .env file found")

	}

}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func SourceAccountInfo() (string, string, string) {
	return os.Getenv("SOURCE_ACCOUNT_NAME"), os.Getenv("SOURCE_ACCOUNT_KEY"), os.Getenv("SOURCE_SHARE_NAME")
}

func DestAccountInfo() (string, string, string) {
	return os.Getenv("DEST_ACCOUNT_NAME"), os.Getenv("DEST_ACCOUNT_KEY"), os.Getenv("DEST_SHARE_NAME")
}
