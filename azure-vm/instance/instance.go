package instance

import (
	"context"
	"fmt"
	"log"

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

func LaunchInstance(ctx context.Context, pp *Parameters) error {
	// generate the tokens
	genSSHK()

	interfaceClient, err := armnetwork.NewInterfacesClient(pp.SubscriptionID, sshk.Token, nil)
	if err != nil {
		return err
	}

	if ok, err := findINetworkNterface(ctx, pp.RG, pp.VmName+pp.VmInterfaceSuffix, interfaceClient); err != nil {
		log.Fatal("finding interface error occured", err)

	} else if ok {
		log.Fatal("the interface already exists in azure : " + pp.VmName + pp.VmInterfaceSuffix + " in resource group " + pp.RG)
	}

	netInterfacePolerResponse, err := interfaceClient.BeginCreateOrUpdate(
		ctx,
		pp.RG,
		pp.VmName+pp.VmInterfaceSuffix,
		armnetwork.Interface{
			Location: to.Ptr(pp.Location),
			Properties: &armnetwork.InterfacePropertiesFormat{

				IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
					{
						Name: to.Ptr(pp.VmName + "privipConfig"),
						Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
							Subnet: &armnetwork.Subnet{
								ID: to.Ptr(pp.SubnetID),
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

	vmClient, err := armcompute.NewVirtualMachinesClient(pp.SubscriptionID, sshk.Token, nil)
	if err != nil {
		return err
	}

	fmt.Println("Checking if the vm exists..")
	// findVirtualMachine checks if the vertual machine already exists
	if ok, err := findVirtualMachine(ctx, pp.RG, pp.VmName, vmClient); err != nil {
		log.Fatal("finding virtual machine error occured", err)

	} else if ok {
		log.Fatalf("the virtual machine already exists in azure : %s  in resource group %s", pp.VmName, pp.RG)
	}

	fmt.Println("Creating the vm" + pp.VmName)

	var imageRef *armcompute.ImageReference

	if !pp.VmParametersGallery.Enabled {

		imageRef = &armcompute.ImageReference{
			Offer:     to.Ptr(pp.Offer),
			Publisher: to.Ptr(pp.Publisher),
			SKU:       to.Ptr(pp.Sku),
			Version:   to.Ptr(pp.Version),
		}
	} else {

		imageRef = &armcompute.ImageReference{
			SharedGalleryImageID: to.Ptr(pp.VmSharedGallery),
		}
	}
	parameters := armcompute.VirtualMachine{
		Location: to.Ptr(pp.Location),
		Identity: &armcompute.VirtualMachineIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeNone),
		},
		Properties: &armcompute.VirtualMachineProperties{
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: imageRef,
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr(pp.VmName + pp.OSDiskSuffix),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(pp.StorageAccountType),
					},
					DiskSizeGB: to.Ptr(pp.DiskSizeGB), // default 127G
				},
			},
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes(pp.VMType)), // VM size include vCPUs,RAM,Data Disks,Temp storage.
			},
			OSProfile: &armcompute.OSProfile{
				ComputerName:  to.Ptr(pp.VmName),
				AdminUsername: to.Ptr(pp.AdminUsername),
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

	pollerResponse, err := vmClient.BeginCreateOrUpdate(ctx, pp.RG, pp.VmName, parameters, nil)
	if err != nil {
		return err
	}

	vmResponse, err := pollerResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Virtual Machine %v is created...\n", *vmResponse.Name)
	}

	return nil

}
