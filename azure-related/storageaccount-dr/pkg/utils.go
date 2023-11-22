package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

func RestoreReplica(ctx context.Context, stClient *armstorage.AccountsClient, storageAccount, resourceGroup string) error {
	var storage armstorage.SKU
	storage.Name = to.Ptr(armstorage.SKUNameStandardRAGRS)
	// sleep a while because azure API might be not ready after fail over

	time.Sleep(time.Minute * 2)

	resp, err := stClient.Update(ctx, resourceGroup, storageAccount, armstorage.AccountUpdateParameters{
		SKU: &storage,
	}, nil)
	if err != nil {
		log.Printf("failed restore replication: %v", err)
		return err

	}

	log.Println("Replication restored, please check on the portal")
	log.Printf("SKU: %v  , Account: %v, Location: %v", *resp.SKU.Name, *resp.Account.Name, *resp.Location)

	return nil

}

func FailoverStorageaccount(ctx context.Context, stClient *armstorage.AccountsClient, storageAccount, resourceGroup string) error {
	st, err := stClient.BeginFailover(ctx, resourceGroup, storageAccount, nil)
	if err != nil {
		log.Printf("failed to create armstorage client: %v", err)
		return err

	}

	storageAccountResponse, err := st.PollUntilDone(ctx, nil)

	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println("Failover is done", storageAccountResponse)
	}

	return nil
}

func Help() {
	fmt.Println("You have to use exactly three parametes")
	fmt.Println("Also first two are mandatory for storage account and resource group aslo the subscription")

	fmt.Println("storagedr --account examplestorageaccount --rg Azure-rg --subscription subsid --recover-replication ## only for restore RAGRS replicateion")
	fmt.Println("storagedr --account examplestorageaccount --rg Azure-rg --subscription subsid --fail-over ## only for fail over of the storage account")
	fmt.Println("storagedr --account examplestorageaccount --rg Azure-rg --subscription subsid --fail-over-and-replicate ## only for fail over and restore of replication to opposite region")

}
