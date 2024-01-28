package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func getSecretClient(vaultBaseURL string) *azsecrets.Client {

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

func displaySecretExpiration(secretName string, keyvault string) (string, error) {

	client := getSecretClient(keyvault)

	secretBundle, err := client.GetSecret(context.TODO(), secretName, "", &azsecrets.GetSecretOptions{})
	if err != nil {
		return "", errors.New("failed to get the secret")
	}
	// Check if the secret has an expiration time
	if secretBundle.Attributes != nil && secretBundle.Attributes.Expires != nil {
		expirationTime := *secretBundle.Attributes.Expires
		t := formatDate(expirationTime)
		return t, nil
	} else {
		return "no expiration time set in KV secret", nil
	}
}

func formatDate(s time.Time) string {
	expinString := s.Format("2006-01-02 15:04:05")
	return expinString
}
