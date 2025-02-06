package main

import (
	"flag"
	"fmt"
	"go-remoteConn/pkg/types"
	"log"
	"os"
)

var filenamePath string

func main() {
	flag.StringVar(&filenamePath, "f", "", "The filename or path with the filename is requred")
	flag.Parse()
	if filenamePath == "" {
		fmt.Println("Error: The filename or path must be provided.")
		flag.PrintDefaults() // Prints the usage message
		os.Exit(1)           // Exit with a non-zero status
	}

	categories, err := types.LoadConnections(filenamePath)
	if err != nil {
		log.Fatalf("Error loading connections: %v", err)
	}

	fmt.Println("Available Connections:")

	currentIndex := 0
	for _, cat := range categories {
		fmt.Printf("Category: %s\n", cat.Name)
		for _, conn := range cat.Connections {
			fmt.Printf("\tConn: %d: %s, (%v)\n", currentIndex, conn.Name, conn.IP)
			currentIndex++
		}
	}
	var choice int
	fmt.Print("Select a connection by number: ")
	_, err = fmt.Scan(&choice)

	if err != nil || choice < 0 || choice > currentIndex-1 {
		log.Fatalf("Invalid choice: %v", err)
	}

	var selectedConn *types.Connection
	currentIndex = 0
	for _, cat := range categories {
		for _, conn := range cat.Connections {
			if currentIndex == choice {
				selectedConn = &conn
				break
			}
			currentIndex++
		}
		if selectedConn != nil {
			break
		}
	}

	if selectedConn == nil {
		log.Fatalf("No connection found for the selected choice.")
	}

	// Connect to the selected server
	if err := selectedConn.ConnectToServer(); err != nil {
		log.Fatalf("Error connecting to the server: %v", err)
	}
}
