package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type SshKeys struct {
	Token      azcore.TokenCredential
	PublicKey  []byte
	PrivateKey []byte
}

func (s *SshKeys) MyGenerateKeys() error {
	k, p, err := GenerateKeys()
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}
	s.PublicKey = p
	s.PrivateKey = k
	if err = os.WriteFile("azureadmin.pem", s.PrivateKey, 0600); err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	if err = os.WriteFile("azureadmin.pub", s.PublicKey, 0644); err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	return nil

}

func GenerateKeys() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}

	pubKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}

	return pem.EncodeToMemory(privateKeyPEM), ssh.MarshalAuthorizedKey(pubKey), nil
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
