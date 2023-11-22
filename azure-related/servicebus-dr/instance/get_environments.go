package instance

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Parameters struct {
	Location                   string
	PRRG                       string
	DRRG                       string
	Context                    context.Context
	SubscriptionID             string
	DisasterRecoveryConfigName string
	PriNamespaceName           string
	SecNamespaceName           string
	ServiceBusTopics           string
}

func GetEnvs() *Parameters {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")

	}
	var pp Parameters

	pp.Location = os.Getenv("AZURE_LOCATION")
	pp.PRRG = os.Getenv("AZURE_PRI_RESOURCEGROUP")
	pp.DRRG = os.Getenv("AZURE_DRI_RESOURCEGROUP")
	pp.DisasterRecoveryConfigName = os.Getenv("AZURE_DR_CONFIG_NAME")
	pp.PriNamespaceName = os.Getenv("AZURE_SERVICEBUS_PRI_NAMESPACE")
	pp.SecNamespaceName = os.Getenv("AZURE_SERVICEBUS_SEC_NAMESPACE")
	pp.SubscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	pp.Context = context.Background()
	pp.ServiceBusTopics = os.Getenv("AZURE_SERVICEBUS_TOPICS")

	if len(pp.Location) == 0 {
		log.Fatal("You must set your 'AZURE_LOCATION' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.PRRG) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.DRRG) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.DisasterRecoveryConfigName) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.PriNamespaceName) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.SecNamespaceName) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.ServiceBusTopics) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.SubscriptionID) == 0 {
		log.Fatal("You must set your 'AZURE_SUBSCRIPTION_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	return &pp
}
