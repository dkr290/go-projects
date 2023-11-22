# Azure vmss application or command line tool to build azure virtual machine scaleset from internal image Gallery

```

#########VMSS related parameters
#Note!

#export on the command line export AZURE_PASSWORD="xxxx" in the env The idea is to be not hardcoded
or in windoes use set AZURE_PASSWORD=xxxx


AZURE_LOCATION="Location"  ## the location for the VMSS
AZURE_RESOURCEGROUP="rg"  ## the resource group for the VMSS
AZURE_VNET_RG="vnet-rg"  ## the virtual network resource group for the network of VMSS nodes
AZURE_VNET_ID="vnet"  ## the subnet name for the VMSS
AZURE_SUBNET_ID="subnet-scaleset-uat"  ## the subnet where vmss nodes will be installed as networking
AZURE_SUBSCRIPTION_ID="xxxxxxxxxxxxxxxxxxxxxxxxxxxxx" ## The subscription ID
AZURE_VMSSNAME="virtual-machine-scaleset-name-vmss"  ## the scaleset name
AZURE_VMSS_SKU="Standard_F4s_v2"  ## the sku for the scaleset
AZURE_USERNAME="azureuser" ## the admin username
AZURE_VMSS_CAPACITY="1"   # capacity of the nodes
AZURE_COMPUTERNAME_VMSS="somenodenameforthenodeinstancesvmss"






AZURE_GALLERY_SUBSCRIPTION_ID="xxxxxxxxxxxxxxxxxxxxxxxxxxxx"  #if not supplies it will be = to the AZURE_SUBSCRIPTION_ID
AZURE_IMAGE_GALLERY="imagegallery name"   ###The image gallery name
AZURE_IMAGE_GALLERY_RG="gallery-rg"  ## the image gallery resource group
AZURE_IMAGE_NAME="image/versions/0.0.1" # the image and the version

```

## Running the executable

```

azure-vmss.exe --help
azure-vmss.exe --envfile  <the env file path>

```


# Building the binary files


# Linux
```
env GOOS=linux GOARCH=amd64 go build -o bin/vmss-linux-0.0.1
```


```
env GOOS=windows GOARCH=amd64 go build -o bin/vmss-windows-0.0.1.exe
```