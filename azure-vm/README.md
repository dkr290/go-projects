##create a file .env with the following content and replace the values

```
AZURE_LOCATION="North Europe, etc"
AZURE_RESOURCEGROUP="<This is the VM resource group>"
AZURE_VNET_RG="<this is the vnet with subnets resource group for the network for VM>"
AZURE_VNET_ID="<vnet id in form of vnet name>"
AZURE_SUBNET_ID="<subnet name for the vnet>"
AZURE_SUBSCRIPTION_ID="<subscription ID>"
AZURE_VMNAME="Virtual machine name like vmname01 etc"

AZURE_ENABLE_MARKETPLACE="false" // enable or disable from marketplace
AZURE_OFFER="Oracle-Linux" //example
AZURE_PUBLISHER="Oracle"   //example
AZURE_SKU="ol91-lvm-gen2"   //example
AZURE_VERSION="latest"


AZURE_ENABLE_IMAGE_GALLERY="true"
AZURE_IMAGE_GALLERY="image gallery name"
AZURE_IMAGE_GALLERY_RG="image gallery resource group"
AZURE_IMAGE_NAME="imagename/versions/0.0.1" //example needs this path to exact verion otherwise if omit will be to the latest

AZURE_VM_DISKSIZE="64"  //example disk size
AZURE_VM_DISKSUFF="disk-01"  //suffix for the disk
AZURE_VM_TYPE="Standard_B2s"  //template
AZURE_VM_INTERFACE_SUFF="interface-01"  //name just some name for the interfface 
AZURE_VM_ADMINUSERNAME="azureadmin"  //admin username for the vm
AZURE_VM_OSSORAGEACCOUNTTYPE="Standard LRS"  //what kind of diks "Standard LRS", "Premium LRS",  "Premium ZRS" ,"Premium LRS V2","Standard SSD LRS","Standard SSD ZRS"

```