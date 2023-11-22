package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	utils "storageaccount-dr/pkg"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func getCreds() *azidentity.DefaultAzureCredential {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}
	return cred
}

func main() {

	flag.String("account", "", "Storage account name")
	flag.String("subscription", "", "Subscription ID")
	flag.String("rg", "", "Resource group name")
	flag.Bool("recover-replication", false, "Recover replication of storage account with RAGRS")
	flag.Bool("fail-over", false, "Fail over the storage account")
	flag.Bool("fail-over-and-replicate", false, "Fail over the storage account")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	RecoverReplication := viper.GetInt("recover-replication") // retrieve value from viper
	FailOverAndReplicate := viper.GetInt("fail-over-and-replicate")
	FailOver := viper.GetInt("fail-over")
	StAccount := viper.GetString("account")
	RG := viper.GetString("rg")
	Subscription := viper.GetString("subscription")

	if StAccount == "" {
		log.Println("storage account parameter is not defined")
		os.Exit(1)
	}

	if RG == "" {
		log.Println("Resource group  is not defined ")
		os.Exit(1)
	}
	if Subscription == "" {
		log.Println("Subscription is not defined ")
		os.Exit(1)
	}

	if len(os.Args) != 8 {

		fmt.Println("help")
		fmt.Println(len(os.Args))
		utils.Help()
		os.Exit(1)
	}

	ctx := context.Background()
	creds := getCreds()

	clientFactory, err := armstorage.NewClientFactory(Subscription, creds, nil)

	if err != nil {
		log.Printf("failed to create armstorage client: %v", err)

	}
	stClient := clientFactory.NewAccountsClient()

	if FailOverAndReplicate == 1 {
		err = utils.FailoverStorageaccount(ctx, stClient, StAccount, RG)
		if err != nil {
			log.Println(err)
			return
		} else {
			utils.RestoreReplica(ctx, stClient, StAccount, RG)
		}
	}

	if RecoverReplication == 1 {
		utils.RestoreReplica(ctx, stClient, StAccount, RG)
	}

	if FailOver == 1 {
		err = utils.FailoverStorageaccount(ctx, stClient, StAccount, RG)
		if err != nil {
			log.Println(err)
			return
		}
	}

}
