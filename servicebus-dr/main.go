package main

import (
	"fmt"
	"log"

	"github.com/dkr290/go-projects/servicebus-dr/instance"
)

func main() {

	parameters := instance.GetEnvs()

	test, err := instance.GetDisasterRecoveryConfig(parameters.Context, parameters)
	if err != nil {
		log.Fatalln("Launch instance error", err)
	}

	fmt.Println("Disaster Recovery Config", *test.ID)
	fmt.Println("Disaster Recovery alias name", *test.Name)
	fmt.Println("Disaster Recovery secondary namespace", *test.Properties.PartnerNamespace)

}
