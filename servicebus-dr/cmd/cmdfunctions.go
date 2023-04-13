package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dkr290/go-projects/servicebus-dr/instance"
)

func FailoverFromOneToOther(ctx context.Context, ResouceGroup, PrimaryServiceBusNamespace, SecondaryServiceBusNamespace, DisasterRecoveryConfigName, SubscriptionID string) {

	fmt.Printf("Do you want to start the failover from %s to %s for alias %s: y\\n :", PrimaryServiceBusNamespace, SecondaryServiceBusNamespace, DisasterRecoveryConfigName)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if text == 'y' {
			fmt.Println("starting the failover...")
			if err := instance.ServiceBusFailover(ctx, ResouceGroup, SecondaryServiceBusNamespace, DisasterRecoveryConfigName, SubscriptionID); err != nil {
				log.Fatalf(err.Error())
			}
			break

		} else if text == 'n' {
			fmt.Println("Abording the failover")
			os.Exit(1)
		} else {
			fmt.Println("You have to specify either y or n y\\n: ")

		}
	}

}
