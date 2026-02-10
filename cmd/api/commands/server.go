package commands

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nikhildev/bugsbunny/cmd/api/routes"
	"github.com/nikhildev/bugsbunny/database"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type serverConfig struct {
	Host string
	Port string
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Start the BugsBunny API server.`,
	Run: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		v.AutomaticEnv()

		v.SetEnvPrefix("HTTP_SERVER")
		v.SetConfigFile(".env")
		v.SetConfigType("env")
		err := v.ReadInConfig()

		if err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}
		serverConfig := serverConfig{
			Host: v.GetString("HTTP_SERVER_HOST"),
			Port: v.GetString("HTTP_SERVER_PORT"),
		}

		// Initialize database
		_, err = database.InitDB(database.GetDbConfig())
		if err != nil {
			log.Fatalf("Error initializing database: %v", err)
		}
		defer database.CloseDbClient()
		// Use migrate command to migrate the database

		fmt.Println("Starting server on", serverConfig.Host, ":", serverConfig.Port)
		mux := routes.SetupRoutes()
		err = http.ListenAndServe(serverConfig.Host+":"+serverConfig.Port, mux)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	},
}
