#Build

### 1. Build Executable per operating system for example Linux amd64

env GOOS=linux GOARCH=amd64 go build -o file

### 2. For windows
env GOOS=windows GOARCH=amd64 go build -o file.exe

# Usage

NOTE: We need environment variables to be set for being able to authenticate
AZURE_CLIENT_ID
AZURE_CLIENT_SECRET
AZURE_TENANT_ID

```
./blobfilecopy --help

 -container string
        The container subfolder where the filename resides to download
  -destfile string
        The destination file name with the path in the local file system like /mnt/file.txt or c:\temp\file1.txt
  -help
        Show help
  -sourcefile string
        Source File in the blob storage
  -storage string
        The storage account name (default "scammodelsdev")
```

### Example
./blobfilecopy -storage scammodelsdev -container constraints -sourcefile constraints.txt -destfile /c/Temp/anotherconstraints.txt