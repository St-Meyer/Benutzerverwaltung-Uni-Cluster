package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	var keyPath string
	if os.Getenv("OS") == "Windows_NT" {
		keyPath = filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_ed25519")
	} else {
		keyPath = filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bitte geben Sie Ihren Benutzernamen ein: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	fmt.Println("Bitte geben Sie die Cluster-Adresse ein: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	fmt.Println("Bitte geben Sie den inner Node an: ")
	innerNode, _ := reader.ReadString('\n')
	innerNode = strings.TrimSpace(innerNode)

	fmt.Println("Bitte geben Sie ihren Befehl an: ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	// Verbindung zum Headnode
	headClient, _, err := headnodeConnection(user, host, keyPath)
	if err != nil {
		log.Fatalf("failed to connect to headnode: %s", err)
	}
	defer headClient.Close()

	fmt.Println("Success to connect to headnode.")

	// Verbindung zum Inner Node
	innerClient, err := connectToInnerNode(headClient, innerNode, keyPath, user)
	if err != nil {
		log.Fatalf("failed to connect to the inner node: %v", err)
	}
	defer innerClient.Close()

	fmt.Println("Success to cnnect to inner node.")

	// Befehl auf dem inneren Node ausführen
	output, err := executeCommand(innerClient, command)
	if err != nil {
		log.Fatalf("failure to run command: %v", err)
	}

	fmt.Printf("Ausgabe: \n%s\n", output)
}
