package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting server...")
		// TODO: Start HTTP server
		return nil
	},
}
