package cmd

import (
	"context"

	"github.com/dwarvesf/devpod-provider-paperspace/pkg/paperspace"
	"github.com/loft-sh/log"
	"github.com/spf13/cobra"
)

// StartCmd holds the cmd flags
type StartCmd struct{}

// NewStartCmd defines a command
func NewStartCmd() *cobra.Command {
	cmd := &StartCmd{}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start an instance",
		RunE: func(_ *cobra.Command, args []string) error {
			paperspaceProvider, err := paperspace.NewProvider(log.Default, false)
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

	return startCmd
}

// Run runs the command logic
func (cmd *StartCmd) Run(
	ctx context.Context,
	paperspaceProvider *paperspace.PaperspaceProvider,
	logs log.Logger,
) error {
	return paperspace.Start(paperspaceProvider)
}
