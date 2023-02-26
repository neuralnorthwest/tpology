package cmd

import "github.com/spf13/cobra"

// Main is the entry point for the itool command.
func Main() error {
	return rootCommand().Execute()
}

func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "itool",
		Short: "ZE Inventory Tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(resourceCommand())
	config.Global.SetupFlags(cmd)
	return cmd
}
