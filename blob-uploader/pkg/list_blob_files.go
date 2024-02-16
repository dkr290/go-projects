package pkg

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// getSecretClient(vaultBaseURL string) - connect to azure and obtain credential for service principal with permissions
/*
export AZURE_TENANT_ID="<active_directory_tenant_id"
export AZURE_CLIENT_ID="<service_principal_appid>"
export AZURE_CLIENT_SECRET="<service_principal_password>"
export STORAGE_NAME = "blob storage NAME"
export CONTAINER_NAME = ""ContainerName

*/
func GetBlobFiles(storageName, containerName string) ([]string, error) {

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	StorageURL := "https://" + storageName + ".blob.core.windows.net"
	client, err := azblob.NewClient(StorageURL, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListBlobsFlatPager(containerName, nil)

	// continue fetching pages until no more remain
	var blobItems []string
	for pager.More() {
		// advance to the next page
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}

		// print the blob names for this page

		for _, blob := range page.Segment.BlobItems {
			blobItems = append(blobItems, *blob.Name)

		}
	}

	return blobItems, nil

}
