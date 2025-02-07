package tcmd

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func Tfinit(tf *tfexec.Terraform) error {
	color.Cyan("Executing tf init")
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)
	if err := tf.Init(context.Background()); err != nil {
		return fmt.Errorf("error running terraform init: %v", err)
	}
	return nil
}
