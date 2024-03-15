package helpers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

// getSecretClient(vaultBaseURL string) - connect to azure and obtain credential for service principal with permissions
/*
export AZURE_TENANT_ID="<active_directory_tenant_id"
export AZURE_CLIENT_ID="<service_principal_appid>"
export AZURE_CLIENT_SECRET="<service_principal_password>"

*/
func GetSecretClient(vaultBaseURL string) *azsecrets.Client {

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

// displaySecretExpiration(secretName string, keyvault string) - connects to the keyvault and return the expiration dates
func DisplaySecretExpiration(secretName string, keyvault string) (string, error) {

	client := GetSecretClient(keyvault)

	secretBundle, err := client.GetSecret(context.TODO(), secretName, "", &azsecrets.GetSecretOptions{})
	if err != nil {
		return "", errors.New("failed to get the secret" + err.Error())
	}
	// Check if the secret has an expiration time
	if secretBundle.Attributes != nil && secretBundle.Attributes.Expires != nil {
		expirationTime := *secretBundle.Attributes.Expires
		t := FormatDate(expirationTime)
		return t, nil
	} else {
		return "no expiration time set in KV secret", nil
	}
}

func FormatDate(s time.Time) string {
	expinString := s.Format("2006-01-02 15:04:05")
	return expinString
}

func ExtractKVName(urlString string) (string, error) {

	// Parse the URL
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
	}

	// Split the host into subdomains
	subdomains := strings.Split(parsedURL.Hostname(), ".")

	// Extract the desired subdomain
	var desiredSubdomain string
	if len(subdomains) > 0 {
		desiredSubdomain = subdomains[0]
	}

	return desiredSubdomain, nil
}
