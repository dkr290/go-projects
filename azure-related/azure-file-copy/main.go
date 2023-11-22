package main

import (
	"azure-file-copy/pkg/fileitems"
	"azure-file-copy/pkg/helpers"
	"flag"
	"fmt"
	"os"

	"github.com/Azure/azure-storage-file-go/azfile"
)

func main() {
	var envfile string
	help := flag.Bool("help", false, "Show help")
	flag.StringVar(&envfile, "envfile", ".env", "The storage account name")

	// Parse the flag
	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	helpers.GetEnvs(envfile)
	sAccountName, sAccountKey, sSharename := helpers.SourceAccountInfo()

	credential, err := azfile.NewSharedKeyCredential(sAccountName, sAccountKey)
	helpers.HandleError(err)

	fl := fileitems.NewFileIems(sAccountName, sSharename, credential)

	directories := fl.GetDirs()
	fmt.Println(directories)

	mMap := make(map[string][]string)
	for _, d := range directories {

		mMap[d] = fl.GetFiles(d)

	}
	fmt.Println(mMap)

}
