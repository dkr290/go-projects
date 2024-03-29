package main

import (
	"log"

	"github.com/dkr290/go-projects/azure-vm/instance"
)

func main() {

	parameters := instance.GetEnvs()

	if err := instance.LaunchInstance(parameters.Context, parameters); err != nil {
		log.Fatalln("Launch instance error", err)
	}

}
