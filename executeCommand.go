package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func executeCommand(client *ssh.Client, command string) (string, error) {

	// Sitzung starten
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %s", err)
	}
	defer session.Close()

	// Befehl ausführen
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %s", err)
	}
	return string(output), nil
}
