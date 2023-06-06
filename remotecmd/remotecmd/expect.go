package remotecmd

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	expect "github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
)

func installExpect(conn *ssh.Client) (err error) {

	// Here we are setting up an io.Writer that will write
	// to our debug strings.Builder{}. If we run into an
	// error, the output of the command will dumpt to STDERR.
	r, w := io.Pipe()
	debug := strings.Builder{}
	debugDone := make(chan struct{})
	go func() {
		io.Copy(&debug, r)
		close(debugDone)
	}()

	defer func() {
		// Wait for our io.Copy() to be done.
		<-debugDone

		// Only log this if we had an error.
		if err != nil {
			log.Printf("expect debug:\n%s", debug.String())
		}
	}()

	e, _, err := expect.SpawnSSH(conn, 5*time.Second, expect.Tee(w))
	if err != nil {
		return err
	}
	defer e.Close()

	var promptRE = regexp.MustCompile(`\$ `)

	_, _, err = e.Expect(promptRE, 10*time.Second)
	if err != nil {
		return fmt.Errorf("did not get shell prompt")
	}

	if err := e.Send("sudo apt-get install expect\n"); err != nil {
		return fmt.Errorf("error on send command: %s", err)
	}

	_, _, ecase, err := e.ExpectSwitchCase(
		[]expect.Caser{
			&expect.Case{
				R: regexp.MustCompile(`Do you want to continue\? \[Y/n\] `),
				T: expect.OK(),
			},
			&expect.Case{
				R: regexp.MustCompile(`is already the newest`),
				T: expect.OK(),
			},
		},
		10*time.Second,
	)
	if err != nil {
		return fmt.Errorf("apt-get install did not send what we expected")
	}

	switch ecase {
	case 0:
		if err := e.Send("Y\n"); err != nil {
			return err
		}
	}

	_, _, err = e.Expect(promptRE, 10*time.Second)
	if err != nil {
		return fmt.Errorf("did not get shell prompt")
	}

	return nil
}
