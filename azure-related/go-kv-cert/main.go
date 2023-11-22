package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

var vaultURI string
var help = flag.Bool("help", false, "Show help")

var vaultKey string

func main() {

	flag.StringVar(&vaultURI, "key-vault-uri", "", "The URI for the keyvault to export the private key string")
	flag.StringVar(&vaultKey, "key-vault-secret", "", "The secret from whitch to export the private key string value")

	// Parse the flag
	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if vaultURI == "" || vaultKey == "" {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(vaultURI)
	fmt.Println(vaultKey)

	keyvaultClient := getKeyVaultClient(vaultURI)

	// set secretVersion empty string ("") to receive the latest
	secret := GetSecret(keyvaultClient, vaultKey, "")

	SplitCert(secret)

}

func getKeyVaultClient(vaultURI string) (client *azsecrets.Client) {

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
	client = azsecrets.NewClient(vaultURI, cred, nil)
	return client

}

func GetSecret(client *azsecrets.Client, secretName string, version string) (s string) {

	// Get a secret. An empty string version gets the latest version of the secret.

	resp, err := client.GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
	}

	return *resp.Value

}

func SplitCert(s string) {

	const (
		begin = "-----BEGIN OPENSSH PRIVATE KEY-----"
		end   = "-----END OPENSSH PRIVATE KEY-----"
	)

	rawString := strings.ReplaceAll(s, begin, "")
	rawString = strings.ReplaceAll(rawString, end, "")

	theString := strings.ReplaceAll(rawString, " ", "\n")

	finalString := begin + theString + end + "\n"

	savePrivateKey(finalString)

}

func savePrivateKey(key string) {

	err := os.WriteFile("./aks_key", []byte(key), 0400)
	if err != nil {
		log.Fatalln("Error writing the file", err)
	}

}
