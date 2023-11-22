#Build

### 1. Build Executable per operating system for example Linux amd64

env GOOS=linux GOARCH=amd64 go build -o storagedr

### 2. For windows
env GOOS=windows GOARCH=amd64 go build -o storagedr.exe

# Usage



```
./storagedr --help

       --account string            Storage account name
      --fail-over                 Fail over the storage account
      --fail-over-and-replicate   Fail over the storage account
      --recover-replication       Recover replication of storage account with RAGRS
      --rg string                 Resource group name




```

### Example
```
 ## only for restore RAGRS replicateion"
storagedr --account examplestorageaccount --rg Azure-rg --subscription subsid --recover-replication

## only for fail over of the storage account"
storagedr --account examplestorageaccount --rg Azure-rg --subscription subsid --fail-over 


## only for fail over and restore of replication to opposite region"
storagedr --account examplestorageaccount --rg Azure-rg --subscription subsid --fail-over-and-replicate


```