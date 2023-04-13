package instance

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
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

func ServiceBusFailover(ctx context.Context, RG string, Namespace string, disasterRecoveryConfigName string, SubscriptionID string) error {

	disasterRecoveryConfigsClient, err := armservicebus.NewDisasterRecoveryConfigsClient(SubscriptionID, getCreds(), nil)
	if err != nil {
		return err
	}

	_, err = disasterRecoveryConfigsClient.FailOver(ctx, RG, Namespace, disasterRecoveryConfigName, nil)
	if err != nil {
		return err
	}

	return nil

}

func GetDisasterRecoveryConfig(ctx context.Context, RG string, Namespace string, DrconfigName string, SubscriptionID string) (*armservicebus.ArmDisasterRecovery, error) {
	disasterRecoveryConfigsClient, err := armservicebus.NewDisasterRecoveryConfigsClient(SubscriptionID, getCreds(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := disasterRecoveryConfigsClient.Get(ctx, RG, Namespace, DrconfigName, nil)
	if err != nil {
		return nil, err
	}

	return &resp.ArmDisasterRecovery, nil
}

func CreateDisasterRecoveryConfig(ctx context.Context, RG string, Namespace string, DrconfigName string, SecondaryNamespace string, SubscriptionID string) (*armservicebus.ArmDisasterRecovery, error) {
	disasterRecoveryConfigsClient, err := armservicebus.NewDisasterRecoveryConfigsClient(SubscriptionID, getCreds(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := disasterRecoveryConfigsClient.CreateOrUpdate(
		ctx,
		RG,
		Namespace,
		DrconfigName,
		armservicebus.ArmDisasterRecovery{
			Properties: &armservicebus.ArmDisasterRecoveryProperties{
				PartnerNamespace: to.Ptr(SecondaryNamespace),
			},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &resp.ArmDisasterRecovery, nil
}

func GetNamespaceID(ctx context.Context, RG string, Namespace string, SubscriptionID string) (*armservicebus.NamespacesClientGetResponse, error) {
	nsClient, err := armservicebus.NewNamespacesClient(SubscriptionID, getCreds(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := nsClient.Get(ctx, RG, Namespace, nil)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
