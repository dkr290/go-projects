package pkg

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

func GetSubnetID(ctx context.Context, VnetName string, SubnetName string, ResourceGroup string, subnetClient *armnetwork.SubnetsClient) *string {

	client, err := subnetClient.Get(ctx, ResourceGroup, VnetName, SubnetName, nil)

	var errResponse *azcore.ResponseError
	if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
		log.Fatal(err)
	}

	return client.ID

}

func FindImage(ctx context.Context, SubscriptionID string, ImageGallery string, ImageGalleryRG string, ImageName string, sshk SshKeys) armcompute.GalleryImagesClientGetResponse {

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
