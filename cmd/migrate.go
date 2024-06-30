package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lostsnow/keqing/cmd/migrate"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrations",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
	},
}

func init() {
	migrateCmd.AddCommand(migrate.HashCmd)
	migrateCmd.AddCommand(migrate.GenerateCmd)
	migrateCmd.AddCommand(migrate.LintCmd)
	migrateCmd.AddCommand(migrate.ApplyCmd)
}
