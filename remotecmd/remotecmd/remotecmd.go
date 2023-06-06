package remotecmd

import (
	"context"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

type OutputString struct {
	Host string
	Out  string
}

func Connect(protocol string, host string, cmd string, config *ssh.ClientConfig) *OutputString {

	conn, err := ssh.Dial(protocol, host, config)
	if err != nil {
		log.Fatalln("error: could not dial to the host: ", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// if err := installExpect(conn); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)
	// }

	out, err := combinedOutput(ctx, conn, cmd)
	if err != nil {
		log.Println("command error: ", err)
	}

	return &OutputString{
		Host: host,
		Out:  out,
	}
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
