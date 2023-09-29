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
	"github.com/spf13/cobra"
)

// InitCmd holds the cmd flags
type InitCmd struct{}

// NewInitCmd defines a init
func NewInitCmd() *cobra.Command {
	cmd := &InitCmd{}
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Init account",
		RunE: func(_ *cobra.Command, args []string) error {
			paperspaceProvider, err := paperspace.NewProvider(log.Default, true)
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

	return initCmd
}

// Run runs the init logic
func (cmd *InitCmd) Run(
	ctx context.Context,
	paperspaceProvider *paperspace.PaperspaceProvider,
	machine *provider.Machine,
	logs log.Logger,
) error {
	sshFolder := paperspaceProvider.Config.SSHFolder
	if strings.Contains(sshFolder, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		sshFolder = strings.Replace(sshFolder, "~", homeDir, 1)
	}

	fmt.Println(sshFolder)
	_, err := ssh.GetPrivateKeyRawBase(sshFolder)
	if err != nil {
		return fmt.Errorf("load private key: %w", err)
	}
	return paperspace.Init(paperspaceProvider)
}
