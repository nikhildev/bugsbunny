package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	URL      string
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Apply database migrations to the BugsBunny database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadDBConfig()
		fmt.Printf("Running migrations... (host=%s db=%s)\n", cfg.Host, cfg.Name)
		// TODO: Run migrations using cfg
		_ = cfg
		return nil
	},
}

func loadDBConfig() dbConfig {
	v := viper.New()
	v.SetEnvPrefix("DB")
	v.AutomaticEnv()
	v.BindEnv("url", "DATABASE_URL") // Allow DATABASE_URL as well as DB_URL

	// Env vars: DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE, DB_URL or DATABASE_URL
	v.SetDefault("host", "localhost")
	v.SetDefault("port", "5432")
	v.SetDefault("sslmode", "disable")

	return dbConfig{
		Host:     v.GetString("host"),
		Port:     v.GetString("port"),
		User:     v.GetString("user"),
		Password: v.GetString("password"),
		Name:     v.GetString("name"),
		SSLMode:  v.GetString("sslmode"),
		URL:      v.GetString("url"),
	}
}
