package instance

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/dkr290/go-devops/azure-instance/keys"
)

var sshk keys.SshKeys

func genSSHK() {

	if err := sshk.MyGenerateKeys(); err != nil {
		log.Fatalln("Error my generation of keys", err)
	}

	if err := sshk.GetToken(); err != nil {
		log.Fatalln("Error generation the token", err)
	}

}

func LaunchInstance(ctx context.Context, resourceGroupName, location, vnetID, subnetID string, subscription_id, VmName string) error {
	// generate the tokens
	genSSHK()

	interfaceClient, err := armnetwork.NewInterfacesClient(subscription_id, sshk.Token, nil)
	if err != nil {
		return err
	}

	netInterfacePolerResponse, err := interfaceClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		VmName+"interface-01",
		armnetwork.Interface{
			Location: to.Ptr(location),
			Properties: &armnetwork.InterfacePropertiesFormat{

				IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
					{
						Name: to.Ptr(VmName + "privipConfig"),
						Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
							Subnet: &armnetwork.Subnet{
								ID: to.Ptr(subnetID),
							},
							PublicIPAddress: nil,
						},
					},
				},

				NetworkSecurityGroup: nil,
			},
		},
		nil,
	)

	if err != nil {
		return err
	}
	netInterfaceResponse, err := netInterfacePolerResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Network Interface %v is creating...\n", *netInterfaceResponse.Name)
	}

	// Create the vm

	fmt.Println("Creating the vm")
	vmClient, err := armcompute.NewVirtualMachinesClient(subscription_id, sshk.Token, nil)
	if err != nil {
		return err
	}

	parameters := armcompute.VirtualMachine{
		Location: to.Ptr(location),
		Identity: &armcompute.VirtualMachineIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeNone),
		},
		Properties: &armcompute.VirtualMachineProperties{
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{

					Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
					Publisher: to.Ptr("canonical"),
					SKU:       to.Ptr("20_04-lts-gen2"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr(VmName + "disk-01"),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS), // OSDisk type Standard/Premium HDD/SSD
					},
					DiskSizeGB: to.Ptr[int32](50), // default 127G
				},
			},
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_B2s")), // VM size include vCPUs,RAM,Data Disks,Temp storage.
			},
			OSProfile: &armcompute.OSProfile{ //
				ComputerName:  to.Ptr(VmName),
				AdminUsername: to.Ptr("azureadmin"),
				LinuxConfiguration: &armcompute.LinuxConfiguration{
					DisablePasswordAuthentication: to.Ptr(true),
					SSH: &armcompute.SSHConfiguration{
						PublicKeys: []*armcompute.SSHPublicKey{
							{
								Path:    to.Ptr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", "azureadmin")),
								KeyData: to.Ptr(string(sshk.PublicKey)),
							},
						},
					},
				},
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: netInterfaceResponse.ID,
					},
				},
			},
		},
	}

	pollerResponse, err := vmClient.BeginCreateOrUpdate(ctx, resourceGroupName, VmName, parameters, nil)
	if err != nil {
		return err
	}

	vmResponse, err := pollerResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Virtual Machine %v is creating...\n", *vmResponse.Name)
	}

	return nil

}

func findVnet(ctx context.Context, rg, vnetName string, vnetClient *armnetwork.VirtualNetworksClient) (bool, error) {

	_, err := vnetClient.Get(ctx, rg, vnetName, nil)
	if err != nil {
		var errResponse *azcore.ResponseError
		if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
			return false, nil
		}
		return false, err

	}
	return true, nil

}
