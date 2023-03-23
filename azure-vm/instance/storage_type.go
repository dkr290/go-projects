package instance

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

func GetStorageAccountType(StType string) armcompute.StorageAccountTypes {

	var storagetypes armcompute.StorageAccountTypes

	switch StType {
	case "Standard LRS":
		storagetypes = armcompute.StorageAccountTypesStandardLRS // OSDisk type Standard/Premium HDD/

	case "Premium LRS":
		storagetypes = armcompute.StorageAccountTypesPremiumLRS

	case "Premium ZRS":
		storagetypes = armcompute.StorageAccountTypesPremiumZRS

	case "Premium LRS V2":
		storagetypes = armcompute.StorageAccountTypesPremiumV2LRS

	case "Standard SSD LRS":
		storagetypes = armcompute.StorageAccountTypesStandardSSDLRS

	case "Standart SSD ZRS":
		storagetypes = armcompute.StorageAccountTypesStandardSSDZRS

	case "Ultra SSD LRS":
		storagetypes = armcompute.StorageAccountTypesUltraSSDLRS

	default:
		storagetypes = armcompute.StorageAccountTypesStandardLRS
	}

	return storagetypes

}
