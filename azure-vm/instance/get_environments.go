package instance

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Parameters struct {
	Location               string
	RG                     string
	Context                context.Context
	VnetID                 string
	SubnetID               string
	SubscriptionID         string
	VmName                 string
	NetworkSecurityGroupID string
}

func GetEnvs() *Parameters {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	pp := Parameters{
		Location:       os.Getenv("AZURE_LOCATION"),
		RG:             os.Getenv("AZURE_RESOURCEGROUP"),
		VnetID:         os.Getenv("AZURE_VNET_ID"),
		SubnetID:       os.Getenv("AZURE_SUBNET_ID"),
		SubscriptionID: os.Getenv("AZURE_SUBSCRIPTION_ID"),
		VmName:         os.Getenv("AZURE_VMNAME"),
	}

	if len(pp.Location) == 0 {
		log.Fatal("You must set your 'AZURE_LOCATION' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.RG) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VnetID) == 0 {
		log.Fatal("You must set your 'AZURE_VNET_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.SubnetID) == 0 {
		log.Fatal("You must set your 'AZURE_SUBNET_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.SubscriptionID) == 0 {
		log.Fatal("You must set your 'AZURE_SUBSCRIPTION_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VmName) == 0 {
		log.Fatal("You must set your 'AZURE_VMNAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	return &Parameters{
		Location:       pp.Location,
		RG:             pp.RG,
		VnetID:         pp.VnetID,
		SubnetID:       pp.SubnetID,
		SubscriptionID: pp.SubscriptionID,
		Context:        context.Background(),
		VmName:         pp.VmName,
	}

}
