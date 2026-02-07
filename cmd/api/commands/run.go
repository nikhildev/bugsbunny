package commands

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a service",
	Long:  `Run various BugsBunny services.`,
}
