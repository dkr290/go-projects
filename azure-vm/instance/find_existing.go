package instance

import (
	"context"
	"errors"

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
