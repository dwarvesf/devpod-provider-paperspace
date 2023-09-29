package paperspace

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/loft-sh/devpod/pkg/ssh"
)

// GetPublicKey returns the public key of the Paperspace provider
func GetPublicKey(paperspaceProvider *PaperspaceProvider) (string, error) {
	publicKeyBase64, err := GetPublicKeyBase(paperspaceProvider)
	if err != nil {
		return "", err
	}
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return "", nil
	}
	return string(publicKeyBytes), nil
}

// GetPublicKeyBase returns the public key of the Paperspace provider in Base64 string
func GetPublicKeyBase(paperspaceProvider *PaperspaceProvider) (string, error) {
	sshFolder, err := getSSHFolder(paperspaceProvider)
	if err != nil {
		return "", err
	}
	publicKey, err := ssh.GetPublicKeyBase(sshFolder)
	if err != nil {
		return "", err
	}

	return publicKey, nil
}

// GetPublicKey returns the public key of the Paperspace provider
func GetPrivateKey(paperspaceProvider *PaperspaceProvider) (string, error) {
	privateKeyBase, err := GetPrivateKeyBase(paperspaceProvider)
	if err != nil {
		return "", err
	}
	return string(privateKeyBase), nil
}

// GetPrivateKeyBase returns the private key of the Paperspace provider in a byte slice
func GetPrivateKeyBase(paperspaceProvider *PaperspaceProvider) ([]byte, error) {
	sshFolder, err := getSSHFolder(paperspaceProvider)
	if err != nil {
		return nil, err
	}
	privateKey, err := ssh.GetPrivateKeyRawBase(sshFolder)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// getSSHFolder returns the SSH folder of the Paperspace provider
func getSSHFolder(paperspaceProvider *PaperspaceProvider) (string, error) {
	sshFolder := paperspaceProvider.Config.SSHFolder
	if strings.Contains(sshFolder, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		sshFolder = strings.Replace(sshFolder, "~", homeDir, 1)
	}

	return sshFolder, nil
}
