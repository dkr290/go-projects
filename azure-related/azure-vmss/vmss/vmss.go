package vmss

import (
	"azure-vmss/pkg"
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

var imageRef *armcompute.ImageReference

func CreateVmss(ctx context.Context, pp *pkg.Parameters, sshk pkg.SshKeys) error {

	subnetClient, err := armnetwork.NewSubnetsClient(pp.SubscriptionID, sshk.Token, nil)
	if err != nil {
		return err
	}

	SubnetID := pkg.GetSubnetID(ctx, pp.VnetID, pp.SubnetID, pp.VnetRG, subnetClient)

	vmssClient, err := armcompute.NewVirtualMachineScaleSetsClient(pp.SubscriptionID, sshk.Token, nil)
	if err != nil {
		return err
	}

	imageId := pkg.FindImage(ctx, pp.GallerySubscriptionID, pp.VmGalleryName, pp.VmGalleryRG, pp.VmGalleryImageName, sshk)
	imageRef = &armcompute.ImageReference{
		ID: imageId.ID,
	}

	vmssDefinition := armcompute.VirtualMachineScaleSet{

		Location: to.Ptr(pp.Location),
		SKU: &armcompute.SKU{
			Name:     to.Ptr(pp.VmssSKU), // Choose the VM size
			Capacity: to.Ptr(int64(pp.VMCapacity)),
		},
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			Overprovision: to.Ptr(false),
			UpgradePolicy: &armcompute.UpgradePolicy{
				Mode: to.Ptr(armcompute.UpgradeModeManual),
				AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
					EnableAutomaticOSUpgrade: to.Ptr(false),
					DisableAutomaticRollback: to.Ptr(false),
				},
			},

			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: imageRef,
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr(pp.ComputerNamePref),
					AdminUsername:      to.Ptr(pp.User),
					AdminPassword:      to.Ptr(pp.Pass),
				},

				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr(pp.VmssName),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								Primary: to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr(pp.VmssName),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Subnet: &armcompute.APIEntityReference{
												ID: SubnetID,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create the VMSS
	pollerResp, err := vmssClient.BeginCreateOrUpdate(context.Background(), pp.RG, pp.VmssName, vmssDefinition, nil)
	if err != nil {
		return err
	}
	log.Printf("Creating virtual machne %v\n", pp.VmssName)

	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	}

	log.Printf("Virtual Machine Scale Set %v created successfully in %v, and in resourcgroup %s!", *resp.VirtualMachineScaleSet.Name, *resp.Location, pp.RG)

	return nil

}
