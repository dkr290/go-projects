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
	VnetRG string
}

type VMParameters struct {
	Offer     string
	Publisher string
	Sku       string
	Version   string
	Enabled   string
}
type VmParametersGallery struct {
	VmGalleryName      string
	VmGalleryRG        string
	VmGalleryImageName string
	PlanName           string
	PlanPublisher      string
	PlanProduct        string
	Enabled            string
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
	pp.VnetRG = os.Getenv("AZURE_VNET_RG")

	pp.VMParameters.Enabled = os.Getenv("AZURE_ENABLE_MARKETPLACE")
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
	pp.VmParametersGallery.Enabled = os.Getenv("AZURE_ENABLE_IMAGE_GALLERY")
	pp.VmParametersGallery.VmGalleryName = os.Getenv("AZURE_IMAGE_GALLERY")
	pp.VmParametersGallery.VmGalleryImageName = os.Getenv("AZURE_IMAGE_NAME")
	pp.VmParametersGallery.VmGalleryRG = os.Getenv("AZURE_IMAGE_GALLERY_RG")
	pp.VmParametersGallery.PlanName = os.Getenv("AZURE_IMAGE_PLAN_NAME")
	pp.VmParametersGallery.PlanProduct = os.Getenv("AZURE_IMAGE_PLAN_PRODUCT")
	pp.VmParametersGallery.PlanPublisher = os.Getenv("AZURE_IMAGE_PLAN_PUBLISHER")

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

	if len(pp.VnetRG) == 0 {
		log.Fatal("You must set your 'AZURE_VNET_RG' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if pp.VMParameters.Enabled == "true" && pp.VmParametersGallery.Enabled == "false" {
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
	} else if pp.VMParameters.Enabled == "false" && pp.VmParametersGallery.Enabled == "true" {

		if len(pp.VmParametersGallery.VmGalleryName) == 0 {
			log.Fatal("You must set your 'AZURE_IMAGE_GALLERY' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
		}
		if len(pp.VmParametersGallery.VmGalleryImageName) == 0 {
			log.Fatal("You must set your 'AZURE_IMAGE_NAME' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
		}
		if len(pp.VmParametersGallery.VmGalleryRG) == 0 {
			log.Fatal("You must set your 'AZURE_IMAGE_GALLERY_RG' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
		}
		if len(pp.VmParametersGallery.PlanName) == 0 {
			log.Fatal("You must set your 'AZURE_IMAGE_PLAN_NAME' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
		}
		if len(pp.VmParametersGallery.PlanProduct) == 0 {
			log.Fatal("You must set your 'AZURE_IMAGE_PLAN_PRODUCT' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
		}
		if len(pp.VmParametersGallery.PlanPublisher) == 0 {
			log.Fatal("You must set your 'AZURE_IMAGE_PLAN_PUBLISHER' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
		}
	} else {
		log.Fatal("Marketplace image should be enabled and Galerry image disabled or vice versa but not both disabled or enabled")
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

	pp.StorageAccountType = GetStorageAccountType(OSStorageAccountType)

	return &pp
}
