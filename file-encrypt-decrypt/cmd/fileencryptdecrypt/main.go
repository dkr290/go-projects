package main

import (
	"flag"

	"github.com/dkr290/go-projects/file-encrypt-decrypt/fileops"
	"github.com/dkr290/go-projects/file-encrypt-decrypt/generatekey"
)

var (
	genKey       string
	fileNameEncr string
	key          string
	fileNameDecr string

	//FlagError = errors.New("please chek all paramenter or --help for help")
)

func main() {

	flag.StringVar(&genKey, "g", "", "Generate symmetric key to store in vault for aes256 encryption, the string can be any string like key")
	flag.StringVar(&fileNameEncr, "e", "", "The full path to the file to encrypt")
	flag.StringVar(&key, "k", "", "The key for encryption")
	flag.StringVar(&fileNameDecr, "d", "", "The full path of the file to decrypt")

	flag.Parse()

	if genKey != "" {
		generatekey.GenerateKey()
		return
	}

	if fileNameEncr != "" && key != "" {
		fileops.EncryptFile(fileNameEncr, key)
		return

	}

	if fileNameDecr != "" && key != "" {
		fileops.DecryptFile(fileNameDecr, key)
		return
	}

}
