package tcmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func TfApply(tf *tfexec.Terraform) error {
	var response string
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	color.Green("Execute terraform apply - y/n or yes/no: ")
	if _, err := fmt.Scanln(&response); err != nil {
		return fmt.Errorf("%v", err)
	}

	if response == "y" || response == "Y" || response == "yes" {
		// Apply the plan (with auto-approve)
		if err := tf.Apply(context.Background()); err != nil {
			log.Fatalf("error running terraform apply: %v", err)
		}
	} else {
		log.Println("Exitting...")
		os.Exit(0)
	}
	return nil
}
