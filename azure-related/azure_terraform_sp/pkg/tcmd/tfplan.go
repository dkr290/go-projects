package tcmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func CustomTfplan(tf *tfexec.Terraform) (bool, error) {
	var response string
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	color.Green("Execute terraform plan - y/n or yes/no: ")
	if _, err := fmt.Scanln(&response); err != nil {
		return false, fmt.Errorf("%v", err)
	}
	if response == "y" || response == "Y" || response == "yes" {
		planOptions := []tfexec.PlanOption{
			tfexec.Out("tf.plan"), // Specify the output file for the plan

		}
		// Example of running terraform plan
		if _, err := tf.Plan(context.Background(), planOptions...); err != nil {
			return false, fmt.Errorf("error running terraform plan: %v", err)
		}
	} else {
		color.Green("Exitting...")
		return true, nil
	}
	return false, nil
}
