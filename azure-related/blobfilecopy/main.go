package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

var (
	storageURL     string
	containerName  string
	destFileName   string
	sourceFileName string
	help           = flag.Bool("help", false, "Show help")
)

func main() {
	fmt.Printf("Downloading...\n")

	flag.StringVar(&storageURL, "storage", "containerdefault", "The storage account name")
	flag.StringVar(&containerName, "container", "", "The container subfolder where the filename resides to download")
	flag.StringVar(&sourceFileName, "sourcefile", "", "Source File in the blob storage")
	flag.StringVar(&destFileName, "destfile", "", `The destination file name with the path in the local file system like /mnt/file.txt or c:\temp\file1.txt`)

	// Parse the flag
	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if storageURL == "" || containerName == "" || sourceFileName == "" || destFileName == "" {
		flag.Usage()
		os.Exit(0)
	}

	// TODO: replace <storage-account-name> with your actual storage account name

	credential, err := azidentity.NewEnvironmentCredential(nil)

	handleError(err)

	storageURL = "https://" + storageURL + ".blob.core.windows.net"
	client, err := azblob.NewClient(storageURL, credential, nil)
	handleError(err)

	file, err := os.Create(destFileName)
	handleError(err)
	defer file.Close()

	_, err = client.DownloadFile(context.TODO(), containerName, sourceFileName, file, nil)
	handleError(err)

}
