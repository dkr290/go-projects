package tcmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ExecuteTerraformBin() error {
	// Capture and print terraform plan output
	cmd := exec.Command("terraform", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("error running terraform plan: %v", err)
	}

	var response string
	fmt.Print("Execute terraform plan - y/n or yes/no: ")
	if _, err := fmt.Scanln(&response); err != nil {
		log.Fatal(err)
	}
	if response == "y" || response == "Y" || response == "yes" {
		// Capture and print terraform plan output
		cmd = exec.Command("terraform", "plan", "-out", "tf.plan")

		// Get standard output and error pipes
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Fatalf("error running terraform plan: %v", err)
		}
	}

	fmt.Println("Display tfplan content")
	// Capture and print terraform plan output
	cmd = exec.Command("terraform", "show", "tf.plan")

	// Get standard output and error pipes
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("error running terraform plan: %v", err)
	}

	fmt.Print("Execute terraform apply - y/n or yes/no: ")
	if _, err := fmt.Scanln(&response); err != nil {
		log.Fatal(err)
	}
	if response == "y" || response == "Y" || response == "yes" {
		// Capture and print terraform plan output
		cmd = exec.Command("terraform", "apply", "--auto-approve", "tf.plan")

		// Get standard output and error pipes
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Fatalf("error running terraform apply: %v", err)
		}
	}
	return nil
}
