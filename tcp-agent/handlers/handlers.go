package handlers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh"
)

func HandleClient(conn net.Conn, sshConfig *ssh.ServerConfig) {
	defer conn.Close()

	// Perform SSH handshake
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, sshConfig)
	if err != nil {
		fmt.Println("Error accepting SSH connection:", err)
		return
	}
	defer sshConn.Close()
	fmt.Println("SSH connection established")

	// Handle global out-of-band requests
	go ssh.DiscardRequests(reqs)

	// Accept channel requests
	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			if err := newChannel.Reject(ssh.UnknownChannelType, "unsupported channel type"); err != nil {
				log.Println(err)
			}
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			fmt.Println("Error accepting channel:", err)
			return
		}
		go func(in <-chan *ssh.Request) {
			for req := range in {
				switch req.Type {
				case "shell":
					if len(req.Payload) == 0 {
						req.Reply(true, nil)
						go shell(channel)
					}
				default:
					req.Reply(false, nil)
				}
			}
		}(requests)
	}
}
func shell(channel ssh.Channel) {
	defer channel.Close()

	// Welcome message
	if _, err := channel.Write([]byte("Welcome to the remote shell\n")); err != nil {
		log.Println(err)
	}

	// Create a scanner to read commands from the client
	scanner := bufio.NewScanner(channel)

	// Display prompt and read commands
	for {
		// Display shell prompt
		if _, err := channel.Write([]byte("$ ")); err != nil {
			log.Println(err)
		}

		// Read command from the client
		if !scanner.Scan() {
			// End of input
			break
		}

		// Get the command entered by the user
		command := scanner.Text()
		if command == "exit" {
			fmt.Println("exit the remote shell")
			return
		}

		// Execute the command and send back the output
		output, err := executeCommand(command)
		if err != nil {
			channel.Write([]byte("Error executing command: " + err.Error() + "\n"))
		} else {
			channel.Write(output)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from shell:", err)
		return
	}
}

func executeCommand(command string) ([]byte, error) {
	// Split command into command and arguments
	parts := strings.Fields(command)
	fmt.Println("Command parts:", parts) // Add this line for debugging
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty command")
	}
	//handle cd command separately
	if parts[0] == "cd" {
		if len(parts) < 2 {
			return nil, fmt.Errorf("missing directory argument for 'cd' command")
		}
		// change the directory
		err := os.Chdir(parts[1])
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	cmd := exec.Command(parts[0], parts[1:]...)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return output, nil
}
