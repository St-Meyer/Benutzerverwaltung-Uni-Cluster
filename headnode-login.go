package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func headnodeConnection(user, host, keyPath string) (*ssh.Client, ssh.Signer, error) {

	// Private key parsen
	signer, err := loadPrivateKey(keyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// SSH-Client-Config
	config := &ssh.ClientConfig{User: user, Auth: []ssh.AuthMethod{
		ssh.PublicKeys(signer),
	}, HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Verbindung herstellen
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial: %s", err)
	}

	return client, signer, nil
}
