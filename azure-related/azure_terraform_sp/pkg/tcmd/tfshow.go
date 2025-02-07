package tcmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func Tfshow(tf *tfexec.Terraform) error {
	// Show the plan output
	tf.SetStdout(nil)
	tf.SetStderr(nil)
	state, err := tf.Show(context.Background())
	if err != nil {
		return fmt.Errorf("error running terraform show: %v", err)
	}
	planFile, err := tf.ShowPlanFileRaw(context.Background(), "tf.plan")
	if err != nil {
		return fmt.Errorf("error running terraform show plan file: %v", err)
	}

	color.Green("Plan file content:")
	color.Green(planFile)
	color.Green("Format state file version:", state.FormatVersion)
	fmt.Println("")
	return nil
}
