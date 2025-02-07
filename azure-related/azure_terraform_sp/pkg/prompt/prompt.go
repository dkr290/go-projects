package prompt

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func Prompt(
	isCache bool,
	clientID, clientSecret, subscriptionID, tenantID, tversion string,
) (string, error) {
	if !isCache {
		fmt.Print("Enter terraform version (like 1.10.5): ")
		if _, err := fmt.Scanln(&tversion); err != nil {
			return "", fmt.Errorf("Error scanln %v", err)
		}
	}
	if !isCache {
		// Prompt for ARM_SUBSCRIPTION_ID
		fmt.Print("Enter ARM_SUBSCRIPTION_ID: ")
		if _, err := fmt.Scanln(&subscriptionID); err != nil {
			return "", fmt.Errorf("Error scanln %v", err)
		}
	}
	os.Setenv("ARM_SUBSCRIPTION_ID", subscriptionID)

	if !isCache {
		// Prompt for ARM_TENANT_ID
		fmt.Print("Enter ARM_TENANT_ID: ")
		if _, err := fmt.Scanln(&tenantID); err != nil {
			return "", fmt.Errorf("Error scanln %v", err)
		}
	}
	os.Setenv("ARM_TENANT_ID", tenantID)

	if !isCache {
		// Prompt for ARM_CLIENT_ID
		fmt.Print("Enter ARM_CLIENT_ID: ")
		if _, err := fmt.Scanln(&clientID); err != nil {
			return "", fmt.Errorf("Error scanln %v", err)
		}
	}
	os.Setenv("ARM_CLIENT_ID", clientID)

	if !isCache {
		// Prompt for ARM_CLIENT_SECRET
		fmt.Print("Enter ARM_CLIENT_SECRET: ")
		if _, err := fmt.Scanln(&clientSecret); err != nil {
			return "", fmt.Errorf("Error scanln %v", err)
		}
	}

	os.Setenv("ARM_CLIENT_SECRET", clientSecret)

	if !isCache {
		color.Blue("Saving to local cache")
		envVars := map[string]string{
			"ARM_SUBSCRIPTION_ID": subscriptionID,
			"ARM_TENANT_ID":       tenantID,
			"ARM_CLIENT_ID":       clientID,
			"ARM_CLIENT_SECRET":   clientSecret,
			"TERRAFORM_VERSION":   tversion,
		}
		if err := os.Mkdir(".cache", os.ModePerm); err != nil {
			return "", err
		}
		file, err := os.Create(".cache/spdata")
		if err != nil {
			log.Fatalf("error creating file: %v", err)
		}
		defer file.Close()
		for k, v := range envVars {
			_, err := file.WriteString(k + ":" + v + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return tversion, nil
}
