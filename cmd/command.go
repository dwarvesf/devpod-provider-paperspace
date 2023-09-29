package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dwarvesf/devpod-provider-paperspace/pkg/paperspace"
	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/devpod/pkg/ssh"
	"github.com/loft-sh/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandCmd holds the cmd flags
type CommandCmd struct{}

// NewCommandCmd defines a command
func NewCommandCmd() *cobra.Command {
	cmd := &CommandCmd{}
	commandCmd := &cobra.Command{
		Use:   "command",
		Short: "Command an instance",
		RunE: func(_ *cobra.Command, args []string) error {
			paperspaceProvider, err := paperspace.NewProvider(log.Default, false)
			if err != nil {
				return err
			}

			return cmd.Run(
				context.Background(),
				paperspaceProvider,
				provider.FromEnvironment(),
				log.Default,
			)
		},
	}

	return commandCmd
}

// Run runs the command logic
func (cmd *CommandCmd) Run(
	ctx context.Context,
	paperspaceProvider *paperspace.PaperspaceProvider,
	machine *provider.Machine,
	logs log.Logger,
) error {
	command := os.Getenv("COMMAND")

	if command == "" {
		return fmt.Errorf("command environment variable is missing")
	}

	sshFolder := paperspaceProvider.Config.SSHFolder
	if strings.Contains(sshFolder, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		sshFolder = strings.Replace(sshFolder, "~", homeDir, 1)
	}
	privateKey, err := ssh.GetPrivateKeyRawBase(sshFolder)
	if err != nil {
		return fmt.Errorf("load private key: %w", err)
	}

	// get instance
	instance, err := paperspace.GetDevpodInstance(paperspaceProvider)
	if err != nil {
		return err
	}

	sshClient, err := ssh.NewSSHClient("paperspace", instance.PublicIPAddress+":22", privateKey)

	if err != nil {
		return errors.Wrap(err, "create ssh client")
	}

	defer sshClient.Close()

	// run command
	return ssh.Run(ctx, sshClient, command, os.Stdin, os.Stdout, os.Stderr)
}
