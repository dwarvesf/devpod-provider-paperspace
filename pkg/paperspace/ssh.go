package paperspace

import (
	"os"
	"strings"

	"github.com/loft-sh/devpod/pkg/ssh"
)

func GetPublicKey(paperspaceProvider *PaperspaceProvider) (string, error) {
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

func GetPrivateKey(paperspaceProvider *PaperspaceProvider) ([]byte, error) {
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
