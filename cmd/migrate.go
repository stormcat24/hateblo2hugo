package cmd

import (
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate hatenablog to hugo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func initMigrateCmd() {
}