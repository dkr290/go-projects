package types

import (
	"encoding/json"
	"os"
)

type Category struct {
	Name        string       `json:"category"`
	Connections []Connection `json:"connections"`
}

type Connection struct {
	Name           string `json:"name"`
	IP             string `json:"ip"`
	Username       string `json:"username"`
	PrivateKeyPath string `json:"privateKeyPath"`
}

var cat []*Category

func LoadConnections(filename string) ([]*Category, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []*Category{}, err
	}

	if err := json.Unmarshal(data, &cat); err != nil {
		return []*Category{}, err
	}

	return cat, nil
}

// func (c *Connection) ConnectToServer_old() error {
// 	// Read the private key file
// 	privateKey, err := os.ReadFile(c.PrivateKeyPath)
// 	if err != nil {
// 		return fmt.Errorf("unable to read private key: %v", err)
// 	}
//
// 	// Parse the private key
// 	signer, err := ssh.ParsePrivateKey(privateKey)
// 	if err != nil {
// 		return fmt.Errorf("unable to parse private key: %v", err)
// 	}
//
// 	// Set up SSH client configuration
// 	config := &ssh.ClientConfig{
// 		User: c.Username,
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Use proper host key verification in production
// 	}
// 	address := fmt.Sprintf("%s:22", c.IP)
//
// 	// Connect to the server
// 	sshClient, err := ssh.Dial("tcp", address, config)
// 	if err != nil {
// 		return fmt.Errorf("failed to dial: %s", err)
// 	}
//
// 	defer sshClient.Close()
// 	// Start an interactive session
// 	session, err := sshClient.NewSession()
// 	if err != nil {
// 		return fmt.Errorf("failed to create session: %s", err)
// 	}
//
// 	defer session.Close()
//
// 	fmt.Printf("Connected to %s (%s)\n", c.Name, c.IP)
// 	// Set up terminal modes
// 	termWidth, termHeight, err := terminal.GetSize(int(os.Stdout.Fd()))
// 	if err != nil {
// 		return fmt.Errorf("failed to get terminal size: %v", err)
// 	}
//
// 	go func() {
// 		for {
// 			termWidth, termHeight, err := terminal.GetSize(int(os.Stdout.Fd()))
// 			if err != nil {
// 				continue
// 			}
// 			session.WindowChange(termHeight, termWidth)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
//
// 	// Request pseudo-terminal
// 	modes := ssh.TerminalModes{
// 		ssh.ECHO:          1,
// 		ssh.TTY_OP_ISPEED: 14400,
// 		ssh.TTY_OP_OSPEED: 14400,
// 		ssh.ICANON:        1,
// 		ssh.IGNCR:         0,
// 		ssh.INLCR:         0,
// 		ssh.ICRNL:         0,
// 	}
//
// 	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
// 	if err != nil {
// 		return fmt.Errorf("failed to set terminal to raw mode: %v", err)
// 	}
// 	defer func() {
// 		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
// 			fmt.Fprintf(os.Stderr, "failed to restore terminal: %v\n", err)
// 		}
// 	}()
// 	if err := session.RequestPty("vt100", termHeight, termWidth, modes); err != nil {
// 		return fmt.Errorf("request for pseudo terminal failed: %s", err)
// 	}
//
// 	// Set up output and input pipes
//
// 	stdout, err := session.StdoutPipe()
// 	if err != nil {
// 		return fmt.Errorf("failed to create stderr pipe: %s", err)
// 	}
//
// 	stderr, err := session.StderrPipe()
// 	if err != nil {
// 		return fmt.Errorf("failed to create stderr pipe: %s", err)
// 	}
// 	go func() {
// 		if _, err := io.Copy(os.Stdout, stdout); err != nil {
// 			fmt.Fprintf(os.Stdout, "stdout copy error: %v\n", err)
// 		}
// 	}()
// 	//	go io.Copy(os.Stdout, stdout)
// 	go func() {
// 		if _, err := io.Copy(os.Stderr, stderr); err != nil {
// 			fmt.Fprintf(os.Stderr, "stderr copy error: %v\n", err)
// 		}
// 	}()
// 	// Set up stdin for sending input from local terminal to remote shell
// 	stdin, err := session.StdinPipe()
// 	if err != nil {
// 		return fmt.Errorf("failed to create stdin pipe: %s", err)
// 	}
// 	go func() {
// 		if _, err := io.Copy(stdin, os.Stdin); err != nil {
// 			fmt.Fprintf(os.Stderr, "stdin copy error: %v\n", err)
// 		}
// 		stdin.Close()
// 	}()
//
// 	// Start the shell
//
// 	if err := session.Shell(); err != nil {
// 		return fmt.Errorf("failed to start shell: %s", err)
// 	}
//
// 	// Wait for the session to finish (this keeps the terminal open)
// 	if err := session.Wait(); err != nil {
// 		return fmt.Errorf("session finished with error: %s", err)
// 	}
//
// 	return nil
// }
