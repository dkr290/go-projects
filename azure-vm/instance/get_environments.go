package instance

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/joho/godotenv"
)

type Parameters struct {
	Location       string
	RG             string
	Context        context.Context
	VnetID         string
	SubnetID       string
	SubscriptionID string
	VmName         string
	VMParameters
	VMsize
	VmParametersGallery
}

type VMParameters struct {
	Offer     string
	Publisher string
	Sku       string
	Version   string
}
type VmParametersGallery struct {
	VmSharedGallery string
	Enabled         bool
}

type VMsize struct {
	DiskSizeGB         int32                          //127
	OSDiskSuffix       string                         //disk-01
	VMType             string                         //"Standard_B2s"
	VmInterfaceSuffix  string                         //interface-01
	AdminUsername      string                         //azureadmin
	StorageAccountType armcompute.StorageAccountTypes //Standard LRS
}

func GetEnvs() *Parameters {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")

	}
	var pp Parameters
	var OSStorageAccountType string

	pp.Location = os.Getenv("AZURE_LOCATION")
	pp.RG = os.Getenv("AZURE_RESOURCEGROUP")
	pp.VnetID = os.Getenv("AZURE_VNET_ID")
	pp.SubnetID = os.Getenv("AZURE_SUBNET_ID")
	pp.SubscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	pp.Context = context.Background()
	pp.VmName = os.Getenv("AZURE_VMNAME")

	pp.Offer = os.Getenv("AZURE_OFFER")
	pp.Publisher = os.Getenv("AZURE_PUBLISHER")
	pp.Sku = os.Getenv("AZURE_SKU")
	pp.Version = os.Getenv("AZURE_VERSION")

	tmp, err := strconv.Atoi(os.Getenv("AZURE_VM_DISKSIZE"))
	if err != nil {
		log.Fatalln("Error in conversion of DiskSizeGB", err)
	}
	pp.DiskSizeGB = int32(tmp)

	pp.OSDiskSuffix = os.Getenv("AZURE_VM_DISKSUFF")
	pp.VMType = os.Getenv("AZURE_VM_TYPE")
	pp.VmInterfaceSuffix = os.Getenv("AZURE_VM_INTERFACE_SUFF")
	pp.AdminUsername = os.Getenv("AZURE_VM_ADMINUSERNAME")

	OSStorageAccountType = os.Getenv("AZURE_VM_OSSORAGEACCOUNTTYPE")
	pp.VmSharedGallery = os.Getenv("AZURE_SHARED_GALLERY_IMAGE_ID")

	if len(pp.Location) == 0 {
		log.Fatal("You must set your 'AZURE_LOCATION' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.RG) == 0 {
		log.Fatal("You must set your 'AZURE_RESOURCEGROUP' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VnetID) == 0 {
		log.Fatal("You must set your 'AZURE_VNET_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.SubnetID) == 0 {
		log.Fatal("You must set your 'AZURE_SUBNET_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.SubscriptionID) == 0 {
		log.Fatal("You must set your 'AZURE_SUBSCRIPTION_ID' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VmName) == 0 {
		log.Fatal("You must set your 'AZURE_VMNAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.Offer) == 0 {
		log.Fatal("You must set your 'AZURE_OFFER' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.Publisher) == 0 {
		log.Fatal("You must set your 'AZURE_PUBLISHER' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.Sku) == 0 {
		log.Fatal("You must set your 'AZURE_SKU' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.Version) == 0 {
		log.Fatal("You must set your 'AZURE_VERSION' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if pp.DiskSizeGB == 0 || pp.DiskSizeGB < 0 {
		log.Fatal("You must set your 'AZURE_VM_DISKSIZE' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.OSDiskSuffix) == 0 {
		log.Fatal("You must set your 'AZURE_VM_DISKSUFF' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VMType) == 0 {
		log.Fatal("You must set your 'AZURE_VM_TYPE' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VmInterfaceSuffix) == 0 {
		log.Fatal("You must set your 'AZURE_VM_INTERFACE_SUFF' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.AdminUsername) == 0 {
		log.Fatal("You must set your 'AZURE_VM_ADMINUSERNAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(OSStorageAccountType) == 0 {
		log.Fatal("You must set your 'AZURE_VM_OSSORAGEACCOUNTTYPE' environmental variable. or the default 'Standard LRS' will be used See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.VmSharedGallery) == 0 || pp.VmSharedGallery == "" {
		pp.VmParametersGallery.Enabled = false
	} else {
		pp.VmParametersGallery.Enabled = true
	}

	pp.StorageAccountType = GetStorageAccountType(OSStorageAccountType)

	return &pp
}
