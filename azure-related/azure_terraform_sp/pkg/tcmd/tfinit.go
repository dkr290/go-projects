package tcmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func Tfinit(tf *tfexec.Terraform) error {
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)
	if err := tf.Init(context.Background()); err != nil {
		return fmt.Errorf("error running terraform init: %v", err)
	}
	return nil
}
