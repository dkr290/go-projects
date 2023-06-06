package remotecmd

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/golang/glog"
	expect "github.com/google/goexpect"
	"github.com/google/goterm/term"
	"golang.org/x/crypto/ssh"
)

var (
	promptRE = regexp.MustCompile(`\$ `)
	timeout  = 10 * time.Hour
)

func Connect(protocol string, host string, cmd string, config *ssh.ClientConfig) {

	conn, err := ssh.Dial(protocol, host, config)

	if err != nil {
		log.Fatalln("error: could not dial to the host: ", err)
	}
	defer conn.Close()

	e, _, err := expect.SpawnSSH(conn, timeout)
	if err != nil {
		glog.Exit(err)
	}
	defer e.Close()

	e.Expect(promptRE, timeout)
	e.Send(cmd + "\n")
	result, _, _ := e.Expect(promptRE, timeout)

	fmt.Println(term.Greenf("Done!\n"))

	fmt.Printf("%s: result:\n %s\n\n", cmd, result)

}
