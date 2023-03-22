package main

import (
	"log"

	"github.com/dkr290/go-projects/azure-vm/instance"
)

func main() {

	parameters := instance.GetEnvs()

	if err := instance.LaunchInstance(
		parameters.Context,
		parameters.RG,
		parameters.Location,
		parameters.VnetID,
		parameters.SubnetID,
		parameters.SubscriptionID,
		parameters.VmName); err != nil {
		log.Fatalln("Launch instance error", err)
	}

}
