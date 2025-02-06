package main

import (
	"log"
	"set_env_vars/pkg/prompt"
	"set_env_vars/pkg/tcmd"
)

func main() {
	err := prompt.Prompt()
	if err != nil {
		log.Fatal(err)
	}
	err = tcmd.ExecuteTerraform()
	if err != nil {
		log.Fatal(err)
	}
}
