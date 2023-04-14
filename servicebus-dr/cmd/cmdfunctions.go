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

func VerifyDisasterRecoveryState(ctx context.Context, ResouceGroup, NameSpace, DisasterRecoveryConfigName, SubscriptionID string) {

	fmt.Printf("Get disasterRecovery configuration state , please check on the portal as well for %s and Servicebusnamespace %s ", ResouceGroup, NameSpace)

	conf, err := instance.GetDisasterRecoveryConfig(ctx, ResouceGroup, NameSpace, DisasterRecoveryConfigName, SubscriptionID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The following config has been dumped....")
	fmt.Println(conf.ID)
	fmt.Println(conf.Name)

}

func VerificationAfterFailOver(CurrentPrimary, PrimaryServiceBusNamespace, DisasterRecoveryConfigName string) {
	fmt.Printf("Please verify in the portal %s and also %s for alias %s\n", PrimaryServiceBusNamespace, CurrentPrimary, DisasterRecoveryConfigName)
	fmt.Println("Please verify that the alias is existing on the current primary namespace", CurrentPrimary)
	fmt.Printf("Deleting of all the topics from  %s namespace to prepare it for replication from the NEW Primary which is %s \n", PrimaryServiceBusNamespace, CurrentPrimary)
	fmt.Println("Do you want to continue... ? y\\n")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if text == 'y' {

			break

		} else if text == 'n' {
			fmt.Println("Abording the failover")
			os.Exit(1)
		} else {
			fmt.Println("You have to specify either y or n y\\n: ")

		}
	}

}
