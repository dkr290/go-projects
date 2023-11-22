package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dkr290/go-projects/servicebus-dr/instance"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	parameters := instance.GetEnvs()

	// using standard library "flag" package
	//Thoose concern only from primary to secondary new to weu for example
	flag.Bool("fail-over-primary-secondary", false, "Failing over for example Primary to Secondary ")
	flag.Bool("clean-topics-primary", false, "Clean topics from previous primary to prepare it for replication")
	flag.Bool("replicate-secondary-primary", false, "Enable replication from secondary example Secondary and to Primary ")
	//This will be opposite steps
	flag.Bool("fail-over-secondary-primary", false, "Failing over for example Secondary , Primary")
	flag.Bool("clean-topics-secondary", false, "Clean topics from previous primary to prepare it for replication")
	flag.Bool("replicate-primary-secondary", false, "Enable replication from original Primary and to Secondary")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	FailOverPprimarySecondary := viper.GetInt("fail-over-primary-secondary") // retrieve value from viper
	FailOverSecondaryPrimary := viper.GetInt("fail-over-secondary-primary")
	CleanTopisPrimary := viper.GetInt("clean-topics-primary")
	CleanTopisSecondary := viper.GetInt("clean-topics-secondary")
	ReplicateSecondaryPrimary := viper.GetInt("replicate-secondary-primary")
	ReplicatePrimarySecondary := viper.GetInt("replicate-primary-secondary")

	if FailOverPprimarySecondary == 1 {
		fmt.Printf("PHASE 1 fail over %s to %s\n", parameters.PriNamespaceName, parameters.SecNamespaceName)
		FailoverFromOneToOther(parameters.Context, parameters.DRRG, parameters.PriNamespaceName, parameters.SecNamespaceName, parameters.DisasterRecoveryConfigName, parameters.SubscriptionID)

	}

	if FailOverSecondaryPrimary == 1 {
		fmt.Printf("PHASE 4 fail over %s to %s\n", parameters.SecNamespaceName, parameters.PriNamespaceName)
		FailoverFromOneToOther(parameters.Context, parameters.PRRG, parameters.SecNamespaceName, parameters.PriNamespaceName, parameters.DisasterRecoveryConfigName, parameters.SubscriptionID)
	}

	if CleanTopisPrimary == 1 {
		fmt.Printf("PHASE 2 clean topics for %s\n", parameters.PriNamespaceName)
		VerificationAfterFailOver(parameters.SecNamespaceName, parameters.PriNamespaceName, parameters.DisasterRecoveryConfigName)
		instance.GetAndDeleteTopic(parameters.Context, parameters.PRRG, parameters.PriNamespaceName, parameters.SubscriptionID, parameters.ServiceBusTopics)
	}
	if CleanTopisSecondary == 1 {
		fmt.Printf("PHASE 5 clean topics for %s\n", parameters.SecNamespaceName)
		VerificationAfterFailOver(parameters.PriNamespaceName, parameters.SecNamespaceName, parameters.DisasterRecoveryConfigName)
		instance.GetAndDeleteTopic(parameters.Context, parameters.DRRG, parameters.SecNamespaceName, parameters.SubscriptionID, parameters.ServiceBusTopics)
	}

	if ReplicateSecondaryPrimary == 1 {
		fmt.Printf("PHASE 3 create replication from %s to %s\n", parameters.SecNamespaceName, parameters.PriNamespaceName)
		PartnerNamespace, err := instance.GetNamespaceID(parameters.Context, parameters.PRRG, parameters.PriNamespaceName, parameters.SubscriptionID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID of the Primary namespace(new secondary))", *PartnerNamespace.ID)
		result, err := instance.CreateDisasterRecoveryConfig(parameters.Context, parameters.DRRG, parameters.SecNamespaceName, parameters.DisasterRecoveryConfigName, *PartnerNamespace.ID, parameters.SubscriptionID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Replication initiated from %s to %s, please check in azure portal when it is finished", parameters.SecNamespaceName, parameters.PriNamespaceName)
		fmt.Println("Resource ", *result.ID)
		fmt.Println("Pending replication count:", *result.Properties.PendingReplicationOperationsCount)
		fmt.Println("Partner Namespace Seconday", *result.Properties.PartnerNamespace)
	}

	if ReplicatePrimarySecondary == 1 {
		fmt.Printf("PHASE 6 create replication from %s to %s\n", parameters.PriNamespaceName, parameters.SecNamespaceName)
		PartnerNamespace, err := instance.GetNamespaceID(parameters.Context, parameters.DRRG, parameters.SecNamespaceName, parameters.SubscriptionID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID of the Primary namespace(new secondary))", *PartnerNamespace.ID)
		result, err := instance.CreateDisasterRecoveryConfig(parameters.Context, parameters.PRRG, parameters.PriNamespaceName, parameters.DisasterRecoveryConfigName, *PartnerNamespace.ID, parameters.SubscriptionID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Replication initiated from %s to %s, please check in azure portal when it is finished", parameters.PriNamespaceName, parameters.SecNamespaceName)
		fmt.Println("Resource ", *result.ID)

	}

}
