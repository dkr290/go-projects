package tcmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func ExecuteTerraform() error {
	// Specify the version you want to install or use the latest
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.10.5")),
	}
	// Install Terraform and get the path to the binary
	execPath, err := installer.Install(context.Background())
	if err != nil {
		return fmt.Errorf("error installing Terraform: %v", err)
	}

	// Set up your working directory for Terraform
	workingDir := "./"

	// Create a new Terraform instance using the installed binary
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return fmt.Errorf("error creating Terraform instance: %v", err)
	}

	// run Terraform commands like Init, Plan, Apply, etc.
	// terraform init
	if err := Tfinit(tf); err != nil {
		return fmt.Errorf("error running terraform init: %v", err)
	}
	// tf plan
	if err := CustomTfplan(tf); err != nil {
		return err
	}
	// Show the plan output
	if err := Tfshow(tf); err != nil {
		return err
	}

	// Apply state
	if err := TfApply(tf); err != nil {
		return err
	}

	var response string
	color.Blue("Remove the tf.plan file - y/n or yes/no: ")
	if _, err := fmt.Scanln(&response); err != nil {
		return fmt.Errorf("%v", err)
	}

	if response == "y" || response == "Y" || response == "yes" {
		os.Remove("tf.plan")
		color.Green("removed tf.plan file")
	}
	return nil
}
