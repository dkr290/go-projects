package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/dkr290/go-projects/remotecmd/remotecmd"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	private  string
	username string
	host     string
	cmd      string
	err      error
	timeout  = time.After(100 * time.Second)
)

var auth ssh.AuthMethod

// TODO better command line arguments
func main() {

	results := make(chan *remotecmd.OutputString, 10)
	flag.StringVar(&private, "i", "default", `path to the private key, if not include -i the default will be home .ssh/id_rsa,  -i "" to prompt for password`)
	flag.StringVar(&username, "u", "", "Username to connect")
	flag.StringVar(&host, "h", "", "Host to connect or ip address [host:port] or [host] or [ip]")
	flag.StringVar(&cmd, "c", "", "Command to execute")
	password := flag.Bool("p", false, "Use for password prompt /either -i or -p but noth both")

	flag.Parse()

	if username == "" || host == "" || cmd == "" {
		flag.Usage()
		return

	}

	hst := strings.Split(host, ",")
	var newhosts []string

	for _, singleHost := range hst {
		_, _, err := net.SplitHostPort(singleHost)
		if err != nil {
			singleHost = singleHost + ":22"
			_, _, err = net.SplitHostPort(singleHost)
			if err != nil {
				log.Fatalln("error: problem with the host passed ", err)
			}
			newhosts = append(newhosts, singleHost)
		}

	}

	if !*password {
		if private == "default" {
			home, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}
			if runtime.GOOS == "windows" {
				private = home + "\\.ssh\\id_rsa"
			} else {
				private = home + "/.ssh/id_rsa"
			}

		}
		auth, err = publicKey(private)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		auth, err = passwordFromTerm()
		if err != nil {
			log.Fatalln(err)
		}
	}

	config := &ssh.ClientConfig{

		User:            username,
		Auth:            []ssh.AuthMethod{auth},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Hour,
	}

	for _, h := range newhosts {

		// result := remotecmd.Connect("tcp", h, cmd, config)
		// fmt.Println(result)

		go func(hostname string) {
			results <- remotecmd.Connect("tcp", hostname, cmd, config)
		}(h)
	}
	for i := 0; i < len(newhosts); i++ {
		select {
		case res := <-results:
			fmt.Println("-----------------------------------------------------------------------------")
			fmt.Printf("Hostname: %s\n", res.Host)
			fmt.Println("")
			fmt.Println(res.Out)
		case <-timeout:
			fmt.Println("Timed out!")
			return
		}
	}

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
