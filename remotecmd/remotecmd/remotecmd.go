package remotecmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func Connect(protocol string, arg []string, config *ssh.ClientConfig) {

	conn, err := ssh.Dial(protocol, arg[1], config)
	if err != nil {
		log.Fatalln("error: could not dial to the host: ", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	out, err := combinedOutput(ctx, conn, arg[2])
	if err != nil {
		log.Println("command error: ", err)
	}
	fmt.Println(out)
}

func combinedOutput(ctx context.Context, conn *ssh.Client, cmd string) (string, error) {
	sess, err := conn.NewSession()
	if err != nil {
		return "", err
	}
	defer sess.Close()

	if v, ok := ctx.Deadline(); ok {
		t := time.NewTimer(v.Sub(time.Now()))
		defer t.Stop()

		go func() {
			x := <-t.C
			if !x.IsZero() {
				sess.Signal(ssh.SIGKILL)
			}
		}()
	}

	b, err := sess.Output(cmd)
	if err != nil {
		return "", err
	}
	return string(b), nil

}
