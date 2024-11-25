package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func loadPrivateKey(keyPath string) (ssh.Signer, error) {
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("failes to parse key: %w", err)
	}

	return signer, nil
}
