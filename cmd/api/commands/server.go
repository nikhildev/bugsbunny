package commands

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nikhildev/bugsbunny/cmd/api/routes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("HTTP_SERVER_PORT")
		if port == "" {
			port = "8080"
		}
		fmt.Printf("Starting server on port %s...\n", port)
		mux := routes.SetupRoutes()
		err := http.ListenAndServe(":"+port, mux)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	},
}

func init() {
	serverCmd.Flags().String("port", "8080", "HTTP server port")
	viper.BindPFlag("HTTP_SERVER_PORT", serverCmd.Flags().Lookup("port"))
	viper.SetEnvPrefix("BUGSBUNNY")
	viper.AutomaticEnv()
}
