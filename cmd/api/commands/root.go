package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bugsbunny",
	Short: "BugsBunny - Bug tracking and management",
	Long:  `BugsBunny is a bug tracking and management application.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(migrateCmd)
}
