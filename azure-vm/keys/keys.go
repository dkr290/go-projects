package keys

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/wardviaene/golang-for-devops-course/ssh-demo"
)

type SshKeys struct {
	Token      azcore.TokenCredential
	PublicKey  []byte
	PrivateKey []byte
}

func (s *SshKeys) MyGenerateKeys() error {
	k, p, err := ssh.GenerateKeys()
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}
	s.PublicKey = p
	s.PrivateKey = k
	if err = os.WriteFile("mykey.pem", s.PrivateKey, 0600); err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	if err = os.WriteFile("mykey.pub", s.PublicKey, 0644); err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	return nil

}

func (s *SshKeys) GetToken() error {

	azCLI, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		fmt.Printf("Token error %v", err)
		return err
	}
	s.Token = azCLI

	return nil
}
