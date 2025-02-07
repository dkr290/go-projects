package tcmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func TfApply(tf *tfexec.Terraform) (bool, error) {
	var response string
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	color.Blue("Execute terraform apply - y/n or yes/no: ")
	if _, err := fmt.Scanln(&response); err != nil {
		return false, fmt.Errorf("%v", err)
	}

	if response == "y" || response == "Y" || response == "yes" {
		// Apply the plan (with auto-approve)
		if err := tf.Apply(context.Background(), tfexec.DirOrPlan("tf.plan")); err != nil {
			return false, fmt.Errorf("error running terraform apply: %v", err)
		}
	} else {
		color.Red("Exitting...")
		return true, nil
	}
	return false, nil
}
