package instance

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func findINetworkNterface(ctx context.Context, ResourceGroup, InterfaceName string, interfaceClient *armnetwork.InterfacesClient) (bool, error) {

	_, err := interfaceClient.Get(ctx, ResourceGroup, InterfaceName, nil)
	if err != nil {
		var errResponse *azcore.ResponseError
		if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
			return false, nil
		}
		return false, err

	}
	return true, nil

}

func findVirtualMachine(ctx context.Context, ResourceGroup, VmName string, vmClient *armcompute.VirtualMachinesClient) (bool, error) {

	_, err := vmClient.Get(ctx, ResourceGroup, VmName, nil)
	if err != nil {
		var errResponse *azcore.ResponseError
		if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
			return false, nil
		}
		return false, err

	}
	return true, nil

}

func findImage(ctx context.Context, SubscriptionID string, ImageGallery string, ImageGalleryRG string, ImageName string) armcompute.GalleryImagesClientGetResponse {

	ImClient, err := armcompute.NewGalleryImagesClient(SubscriptionID, sshk.Token, nil)
	var errResponse *azcore.ResponseError
	if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
		log.Fatal(err)
	}

	resp, err := ImClient.Get(ctx, ImageGalleryRG, ImageGallery, ImageName, nil)
	if err != nil {
		log.Fatal(err)
	}

	return resp

}

func getSubnetID(ctx context.Context, VnetName string, SubnetName string, ResourceGroup string, subnetClient *armnetwork.SubnetsClient) *string {

	client, err := subnetClient.Get(ctx, ResourceGroup, VnetName, SubnetName, nil)

	var errResponse *azcore.ResponseError
	if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
		log.Fatal(err)
	}

	return client.ID

}
