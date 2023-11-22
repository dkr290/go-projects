package fileitems

import (
	"azure-file-copy/pkg/helpers"
	"context"
	"fmt"
	"net/url"

	"github.com/Azure/azure-storage-file-go/azfile"
)

type FileItems struct {
	Sharename   string
	AccountName string
	Credentials *azfile.SharedKeyCredential
}

func NewFileIems(a string, s string, c *azfile.SharedKeyCredential) *FileItems {
	return &FileItems{
		AccountName: a,
		Sharename:   s,
		Credentials: c,
	}
}

func (f *FileItems) GetDirs() []string {
	u, _ := url.Parse(fmt.Sprintf("https://%s.file.core.windows.net/"+f.Sharename, f.AccountName))
	dirUrl := azfile.NewDirectoryURL(*u, azfile.NewPipeline(f.Credentials, azfile.PipelineOptions{}))
	direResp, err := dirUrl.ListFilesAndDirectoriesSegment(context.Background(), azfile.Marker{}, azfile.ListFilesAndDirectoriesOptions{})
	helpers.HandleError(err)
	directories := getFolders(direResp.DirectoryItems)
	return directories

}

func (f *FileItems) GetFiles(d string) []string {

	var err error

	u, _ := url.Parse(fmt.Sprintf("https://%s.file.core.windows.net/"+f.Sharename+"/"+d, f.AccountName))
	dirUrl := azfile.NewDirectoryURL(*u, azfile.NewPipeline(f.Credentials, azfile.PipelineOptions{}))
	listResp, err := dirUrl.ListFilesAndDirectoriesSegment(context.Background(), azfile.Marker{}, azfile.ListFilesAndDirectoriesOptions{})
	helpers.HandleError(err)

	files := getFiles(listResp.FileItems)
	return files

}

func getFolders(items []azfile.DirectoryItem) []string {
	var sharedSubfolders []string
	for _, dName := range items {
		sharedSubfolders = append(sharedSubfolders, dName.Name)
	}

	return sharedSubfolders
}

func getFiles(items []azfile.FileItem) []string {
	var files []string
	for _, dName := range items {
		files = append(files, dName.Name)
	}

	return files
}
