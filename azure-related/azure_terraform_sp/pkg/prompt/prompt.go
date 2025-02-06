package prompt

import (
	"fmt"
	"os"
)

func Prompt() error {
	var subscriptionID, clientSecret, tenantID, clientID string

	// Prompt for ARM_SUBSCRIPTION_ID
	fmt.Print("Enter ARM_SUBSCRIPTION_ID: ")
	if _, err := fmt.Scanln(&subscriptionID); err != nil {
		return fmt.Errorf("Error scanln %v", err)
	}
	os.Setenv("ARM_SUBSCRIPTION_ID", subscriptionID)

	// Prompt for ARM_TENANT_ID
	fmt.Print("Enter ARM_TENANT_ID: ")
	if _, err := fmt.Scanln(&tenantID); err != nil {
		return fmt.Errorf("Error scanln %v", err)
	}
	os.Setenv("ARM_TENANT_ID", tenantID)

	// Prompt for ARM_CLIENT_ID
	fmt.Print("Enter ARM_CLIENT_ID: ")
	if _, err := fmt.Scanln(&clientID); err != nil {
		return fmt.Errorf("Error scanln %v", err)
	}
	os.Setenv("ARM_CLIENT_ID", clientID)

	// Prompt for ARM_CLIENT_SECRET
	fmt.Print("Enter ARM_CLIENT_SECRET: ")
	if _, err := fmt.Scanln(&clientSecret); err != nil {
		return fmt.Errorf("Error scanln %v", err)
	}
	os.Setenv("ARM_CLIENT_SECRET", clientSecret)

	return nil
}
