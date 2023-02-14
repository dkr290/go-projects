package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dkr290/go-projects/remotecmd/remotecmd"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	private = flag.String("private", "", "The path to the ssh key for this connection")
)

var auth ssh.AuthMethod

// TODO better command line arguments
func main() {

	flag.Parse()
	if len(os.Args) != 4 {
		log.Fatalln("error: command must have 2 args, [host] [command] [user]")
	}

	_, _, err := net.SplitHostPort(os.Args[1])
	if err != nil {
		os.Args[1] = os.Args[1] + ":22"
		_, _, err = net.SplitHostPort(os.Args[1])
		if err != nil {
			log.Fatalln("error: problem with the host passed ", err)
		}
	}

	if *private == "" {
		fi, _ := os.Stdin.Stat()
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			log.Fatal("-private not set, cannot use password when STDIN as a pipe")
		}
		auth, err = passwordFromTerm()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		auth, err = publicKey(*private)
		if err != nil {
			log.Fatalln(err)
		}
	}

	u := os.Args[3]

	config := &ssh.ClientConfig{

		User:            u,
		Auth:            []ssh.AuthMethod{auth},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	remotecmd.Connect("tcp", os.Args, config)

}

func passwordFromTerm() (ssh.AuthMethod, error) {

	fmt.Printf("SSH Password: ")
	p, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}
	fmt.Println("")
	if len(bytes.TrimSpace(p)) == 0 {
		return nil, fmt.Errorf("password was empty string")
	}
	return ssh.Password(string(p)), nil
}

func publicKey(privateKeyFile string) (ssh.AuthMethod, error) {

	k, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(k)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signer), nil
}
