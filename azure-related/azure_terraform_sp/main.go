package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"set_env_vars/pkg/prompt"
	"set_env_vars/pkg/tcmd"
	"strings"
)

// Config struct to hold key-value pairs
type Config struct {
	ARMClientID       string
	ARMClientSecret   string
	TerraformVersion  string
	ARMSubscriptionID string
	ARMTenantID       string
}

// ReadConfig reads the file and parses key-value pairs into a struct
func ReadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2) // Split only at the first occurrence of ":"
		if len(parts) != 2 {
			continue // Skip invalid lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Map the values to struct fields
		switch key {
		case "ARM_CLIENT_ID":
			config.ARMClientID = value
		case "ARM_CLIENT_SECRET":
			config.ARMClientSecret = value
		case "TERRAFORM_VERSION":
			config.TerraformVersion = value
		case "ARM_SUBSCRIPTION_ID":
			config.ARMSubscriptionID = value
		case "ARM_TENANT_ID":
			config.ARMTenantID = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	var tfver string
	if _, err := os.Stat(".cache/spdata"); err == nil {
		config, err := ReadConfig(".cache/spdata")
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		tfver, err = prompt.Prompt(
			true,
			config.ARMClientID,
			config.ARMClientSecret,
			config.ARMSubscriptionID,
			config.ARMTenantID,
			config.TerraformVersion,
		)
		if err != nil {
			log.Fatal(err)
		}

	} else if os.IsNotExist(err) {

		tfver, err = prompt.Prompt(false, "", "", "", "", "")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Some other error occurred (e.g., permission issue)
		fmt.Printf("Error checking file: %v\n", err)
	}
	err := tcmd.ExecuteTerraform(tfver)
	if err != nil {
		log.Fatal(err)
	}
}
