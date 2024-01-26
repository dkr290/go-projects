package keyvaultsecrets

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

var (
	vaultBaseURL = "https://kv.vault.azure.net/"
)

func getSecretClient() *azsecrets.Client {

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	client, err := azsecrets.NewClient(vaultBaseURL, cred, nil)
	if err != nil {
		log.Fatalf("Failed to get Azure credentials: %v", err)
	}

	return client
}

func displaySecretExpiration() {
	secretName := "sp-client-secret"

	client := getSecretClient()

	secretBundle, err := client.GetSecret(context.TODO(), secretName, "", &azsecrets.GetSecretOptions{})
	if err != nil {
		log.Fatalf("Failed to get secret: %v", err)
	}
	// Check if the secret has an expiration time
	if secretBundle.Attributes != nil && secretBundle.Attributes.Expires != nil {
		expirationTime := *secretBundle.Attributes.Expires
		fmt.Printf("Secret '%s' expiration time: %v\n", secretName, expirationTime)
	} else {
		fmt.Printf("Secret '%s' does not have an expiration time.\n", secretName)
	}
}
