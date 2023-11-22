package pkg

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Parameters struct {
	Location         string
	RG               string
	Context          context.Context
	VnetID           string
	SubnetID         string
	SubscriptionID   string
	VmssName         string
	ComputerNamePref string
	User             string
	Pass             string
	VMCapacity       int64
	VmParametersGallery
	VnetRG  string
	VmssSKU string
}

type VmParametersGallery struct {
	VmGalleryName         string
	VmGalleryRG           string
	VmGalleryImageName    string
	GallerySubscriptionID string
}

func GetEnvs(file string) (*Parameters, error) {
	var err error
	if err = godotenv.Load(file); err != nil {
		log.Println("No .env file found")
		return nil, err

	}
	var pp Parameters

	pp.Location = os.Getenv("AZURE_LOCATION")
	pp.RG = os.Getenv("AZURE_RESOURCEGROUP")
	pp.VnetID = os.Getenv("AZURE_VNET_ID")
	pp.SubnetID = os.Getenv("AZURE_SUBNET_ID")
	pp.SubscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	pp.Context = context.Background()
	pp.VmssName = os.Getenv("AZURE_VMSSNAME")
	pp.VmssSKU = os.Getenv("AZURE_VMSS_SKU")
	pp.User = os.Getenv("AZURE_USERNAME")
	pp.Context = context.Background()
	pp.Pass = os.Getenv("AZURE_PASSWORD")

	pp.VMCapacity, err = strconv.ParseInt(os.Getenv("AZURE_VMSS_CAPACITY"), 10, 64)
	if err != nil {
		log.Fatal("Cannot Convert capacity for the VMSS to int")
	}
	pp.VnetRG = os.Getenv("AZURE_VNET_RG")
	pp.ComputerNamePref = os.Getenv("AZURE_COMPUTERNAME_VMSS")

	pp.GallerySubscriptionID = os.Getenv("AZURE_GALLERY_SUBSCRIPTION_ID")

	pp.VmParametersGallery.VmGalleryName = os.Getenv("AZURE_IMAGE_GALLERY")
	pp.VmParametersGallery.VmGalleryImageName = os.Getenv("AZURE_IMAGE_NAME")
	pp.VmParametersGallery.VmGalleryRG = os.Getenv("AZURE_IMAGE_GALLERY_RG")

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
	if len(pp.VmssName) == 0 {
		log.Fatal("You must set your 'AZURE_VMNAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.VmssSKU) == 0 {
		log.Fatal("You must set your 'AZURE_VMSS_SKU' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.VnetRG) == 0 {
		log.Fatal("You must set your 'AZURE_VNET_RG' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.User) == 0 {
		log.Fatal("You must set your 'AZURE_USERNAME' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.Pass) == 0 {
		log.Fatal("You must set your 'AZURE_PASSWORD' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if pp.VMCapacity == 0 {
		log.Fatal("You must set your 'AZURE_VMSS_CAPACITY' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.ComputerNamePref) == 0 {
		log.Fatal("You must set your 'AZURE_COMPUTERNAME_VMSS' environmental variable. See\n\t https://pkg.go.dev/os#Getenv")
	}

	if len(pp.VmParametersGallery.VmGalleryName) == 0 {
		log.Fatal("You must set your 'AZURE_IMAGE_GALLERY' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VmParametersGallery.VmGalleryImageName) == 0 {
		log.Fatal("You must set your 'AZURE_IMAGE_NAME' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.VmParametersGallery.VmGalleryRG) == 0 {
		log.Fatal("You must set your 'AZURE_IMAGE_GALLERY_RG' environmental variable.  See\n\t https://pkg.go.dev/os#Getenv")
	}
	if len(pp.GallerySubscriptionID) == 0 {
		pp.GallerySubscriptionID = pp.SubscriptionID
	}

	return &pp, nil
}
