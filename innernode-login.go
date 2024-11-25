package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func connectToInnerNode(proxyClient *ssh.Client, nodeAdress, keyPath, user string) (*ssh.Client, error) {

	// Private key parsen
	signer, err := loadPrivateKey(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// SSH-Client-Config für den inneren Node
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Erstelle einen Tunnel zum inneren Node
	conn, err := proxyClient.Dial("tcp", nodeAdress+":22")
	if err != nil {
		return nil, fmt.Errorf("failed to dial to inner node: %w", err)
	}

	// Aufbau der Connection
	innerClient, chans, reqs, err := ssh.NewClientConn(conn, nodeAdress+":22", config)
	if err != nil {
		return nil, fmt.Errorf("failed to create a inner Connection: %w", err)
	}
	client := ssh.NewClient(innerClient, chans, reqs)
	return client, nil
}
