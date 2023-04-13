package main

import (
	"flag"
	"fmt"

	"github.com/dkr290/go-projects/servicebus-dr/instance"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	parameters := instance.GetEnvs()

	// using standard library "flag" package
	flag.Bool("fail-over-primary-secondary", false, "Failing over for example Primary is North Europe, Secondary is West Europe")
	flag.Bool("fail-over-secondary-primary", false, "Failing over for example Primary is North Europe, Secondary is West Europe")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	FailOverPprimarySecondary := viper.GetInt("fail-over-primary-secondary") // retrieve value from viper
	FailOverSecondaryPrimary := viper.GetInt("fail-over-secondary-primary")
	fmt.Println(FailOverPprimarySecondary)

	if FailOverPprimarySecondary == 1 {
		FailoverFromOneToOther(parameters.Context, parameters.DRRG, parameters.PriNamespaceName, parameters.SecNamespaceName, parameters.DisasterRecoveryConfigName, parameters.SubscriptionID)
	}

	if FailOverSecondaryPrimary == 1 {
		fmt.Println("sec-pri")
	}

	// //WE to NEU to wrap in separate fuctions

	// PartnerNamespace, err := instance.GetNamespaceID(parameters.Context, parameters.RG, parameters.PriNamespaceName, parameters.SubscriptionID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*PartnerNamespace.ID)
	// result, err := instance.CreateDisasterRecoveryConfig(parameters.Context, parameters.DRRG, parameters.SecNamespaceName, parameters.DisasterRecoveryConfigName, *PartnerNamespace.ID, parameters.SubscriptionID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Replication initiated from %s to %s, please check in azure portal wehn it is finished", parameters.SecNamespaceName, parameters.PriNamespaceName)
	// fmt.Println("Resource ", result.ID)
	// fmt.Println("Pending replication count:", result.Properties.PendingReplicationOperationsCount)
	// fmt.Println("Partner Namespace Seconday", result.Properties.PartnerNamespace)

	// fmt.Printf("Do you want to start the failover from %s to %s for alias %s: y\\n :", parameters.SecNamespaceName, parameters.PriNamespaceName, parameters.DisasterRecoveryConfigName)
	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	text, _, err := reader.ReadRune()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if text == 'y' {
	// 		fmt.Println("starting the failover...")
	// 		if err := instance.ServiceBusFailover(parameters.Context, parameters.RG, parameters.PriNamespaceName, parameters.DisasterRecoveryConfigName, parameters.SubscriptionID); err != nil {
	// 			log.Fatalf(err.Error())
	// 		}
	// 		break

	// 	} else if text == 'n' {
	// 		fmt.Println("Abording the failover")
	// 		os.Exit(1)
	// 	} else {
	// 		fmt.Println("You have to specify either y or n y\\n: ")

	// 	}
	// }

	//from neu to WEU replication creation

	//instance.GetAndDeleteTopic(parameters.Context, parameters.DRRG, parameters.SecNamespaceName, parameters.SubscriptionID, parameters.ServiceBusTopics)
	// PartnerNamespace, err := instance.GetNamespaceID(parameters.Context, parameters.DRRG, parameters.SecNamespaceName, parameters.SubscriptionID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*PartnerNamespace.ID)
	// result, err := instance.CreateDisasterRecoveryConfig(parameters.Context, parameters.RG, parameters.PriNamespaceName, parameters.DisasterRecoveryConfigName, *PartnerNamespace.ID, parameters.SubscriptionID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Replication initiated from %s to %s, please check in azure portal wehn it is finished\n", parameters.SecNamespaceName, parameters.PriNamespaceName)
	// fmt.Println("Resource ", *result.ID)
	// fmt.Println("Pending replication count:", *result.Properties.PendingReplicationOperationsCount)
	// fmt.Println("Partner Namespace Seconday", *result.Properties.PartnerNamespace)
}
