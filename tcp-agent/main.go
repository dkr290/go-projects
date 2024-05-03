package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"tcp-agent/handlers"

	"golang.org/x/crypto/ssh"
)

// to connect to the server ssh -i ~/.ssh/id_rsa user@ip -p 8888
func loadRsaPrivateKey() []byte {
	bytes, err := os.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	if err != nil {
		log.Fatalln("cannot read the private key file", err)
	}
	return bytes
}
func main() {
	// Load SSH private key
	privateKeyBytes := loadRsaPrivateKey()
	privateKey, err := ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		fmt.Println("Failed to parse private key:", err)
		return
	}

	// Configure SSH server
	sshConfig := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			return nil, fmt.Errorf("password authentication not supported")
		},
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			// Add your logic here to validate the public key
			// For simplicity, we accept any public key
			return &ssh.Permissions{}, nil
		},
	}

	// Add private key to SSH server configuration
	sshConfig.AddHostKey(privateKey)

	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started, listening on port 8888")

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		// Handle client in a separate goroutine
		go handlers.HandleClient(conn, sshConfig)
	}
}
