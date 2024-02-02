package cmd

import (
	"context"
	"fmt"

	"github.com/dwarvesf/devpod-provider-paperspace/pkg/paperspace"
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
	logs log.Logger,
) error {
	publicKey, err := paperspace.GetPublicKey(paperspaceProvider)
	if err != nil {
		return fmt.Errorf("load private key: %w", err)
	}
	fmt.Printf(
		"Generated SSH key %s/id_devpod_rsa: %s",
		paperspaceProvider.Config.SSHFolder,
		publicKey,
	)

	return paperspace.Init(paperspaceProvider)
}
