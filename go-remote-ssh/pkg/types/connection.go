package types

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
)

func (c *Connection) ConnectToServer() error {
	// Read the private key file
	privateKey, err := os.ReadFile(c.PrivateKeyPath)
	if err != nil {
		return fmt.Errorf("unable to read private key: %v", err)
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("unable to parse private key: %v", err)
	}

	// Set up SSH client configuration
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Use proper host key verification in production
	}
	address := fmt.Sprintf("%s:22", c.IP)

	// Connect to the server
	sshClient, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return fmt.Errorf("failed to dial: %s", err)
	}

	defer sshClient.Close()
	// Start an interactive session
	session, err := sshClient.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}

	defer session.Close()

	// Request pseudo-terminal
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // Enable echo
		ssh.TTY_OP_ISPEED: 14400, // Input speed
		ssh.TTY_OP_OSPEED: 14400, // Output speed
		ssh.ICANON:        1,     // Canonical mode (allows input buffering)
		ssh.ISIG:          1,     // Enable signal processing (Ctrl+C, etc.)
	}

	termWidth, termHeight, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return fmt.Errorf("failed to get terminal size: %v", err)
	}
	// Monitor terminal resize and propagate changes
	go func() {
		for {
			termWidth, termHeight, err := terminal.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				continue
			}
			_ = session.WindowChange(termHeight, termWidth)
			time.Sleep(1 * time.Second)
		}
	}()
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to set terminal to raw mode: %v", err)
	}
	defer func() {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			fmt.Fprintf(os.Stderr, "failed to restore terminal: %v\n", err)
		}
	}()

	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}
	// set input and output
	session.Setenv("TERM", "xterm-256color")
	session.Setenv("LC_CTYPE", "en_US.UTF-8")
	session.Setenv("BASH_ENV", "~/.bashrc")

	// Set up output and input pipes for interactive use
	stdout, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %s", err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %s", err)
	}

	go func() {
		if _, err := io.Copy(os.Stdout, stdout); err != nil {
			fmt.Fprintf(os.Stdout, "stdout copy error: %v\n", err)
		}
	}()

	go func() {
		if _, err := io.Copy(os.Stderr, stderr); err != nil {
			fmt.Fprintf(os.Stderr, "stderr copy error: %v\n", err)
		}
	}()

	// Set up stdin for sending input from local terminal to remote shell
	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdin pipe: %s", err)
	}

	go func() {
		if _, err := io.Copy(stdin, os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "stdin copy error: %v\n", err)
		}
		stdin.Close()
	}()
	// Start the shell
	if err := session.Shell(); err != nil {
		return fmt.Errorf("failed to start shell: %s", err)
	}

	// Wait for the session to finish (this keeps the terminal open)
	if err := session.Wait(); err != nil {
		return fmt.Errorf("session finished with error: %s", err)
	}

	return nil
}
