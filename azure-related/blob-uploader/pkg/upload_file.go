package pkg

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func UploadFileBlob(storageName, containerName, sourcefile string) error {

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		return errors.New("error get credentials" + err.Error())
	}

	StorageURL := "https://" + storageName + ".blob.core.windows.net"
	client, err := azblob.NewClient(StorageURL, cred, nil)
	if err != nil {
		return errors.New("error new client" + err.Error())
	}

	file, err := os.Open(sourcefile)
	if err != nil {
		return errors.New("error opening source file" + err.Error())
	}
	defer file.Close()

	_, err = client.UploadFile(context.Background(), containerName, sourcefile, file, nil)
	if err != nil {
		return err
	}

	return nil

}

func DeleteLocalFiles(sourcefile string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return errors.New("error getting local cwd" + err.Error())

	}
	// Attempt to delete the file
	fullPath := filepath.Join(currentDir, sourcefile)
	err = os.Remove(fullPath)
	if err != nil {
		return errors.New("error deleting local file" + fullPath + err.Error())

	}

	return nil
}
