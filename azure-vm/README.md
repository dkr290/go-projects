##create a file .env with the following content and replace the values


AZURE_LOCATION= <location>
AZURE_RESOURCEGROUP= <azure resource group>
AZURE_VNET_ID= <vnet name >
AZURE_SUBNET_ID= <the subnet id>  //subscriptions/<ID>/resourceGroups/<RG>/providers/Microsoft.Network/virtualNetworks/<VNET>/subnets/<SUBNET>
AZURE_SUBSCRIPTION_ID= <the subscription id> 
AZURE_VMNAME= <virtual machine name>

AZURE_OFFER=<offer> // like 0001-com-ubuntu-server-focal
AZURE_PUBLISHER=<publisher> //like canonical
AZURE_SKU=<sku> like 20_04-lts-gen2
AZURE_VERSION="latest"

AZURE_VM_DISKSIZE="127" // disk size
AZURE_VM_DISKSUFF="disk-suff-01" //disk suffix
AZURE_VM_TYPE="Standard_B2s" //VM type 
AZURE_VM_INTERFACE_SUFF="interface-01" // interface suffix
AZURE_VM_ADMINUSERNAME="admin" //admin user
AZURE_VM_OSSORAGEACCOUNTTYPE="Premium LRS" // account type