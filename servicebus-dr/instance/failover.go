package instance

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

func getCreds() *azidentity.DefaultAzureCredential {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}
	return cred
}

//func ServiceBusFailover(ctx context.Context, pp *Parameters) error {

//}

func GetDisasterRecoveryConfig(ctx context.Context, pp *Parameters) (*armservicebus.ArmDisasterRecovery, error) {
	disasterRecoveryConfigsClient, err := armservicebus.NewDisasterRecoveryConfigsClient(pp.SubscriptionID, getCreds(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := disasterRecoveryConfigsClient.Get(ctx, pp.RG, pp.PriNamespaceName, pp.DisasterRecoveryConfigName, nil)
	if err != nil {
		return nil, err
	}

	return &resp.ArmDisasterRecovery, nil
}
