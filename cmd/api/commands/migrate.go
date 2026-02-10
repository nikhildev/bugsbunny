package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nikhildev/bugsbunny/models"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

var autopopulate bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Apply database migrations to the BugsBunny database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := loadDBConfig()
		fmt.Printf("Running migrations... (host=%s db=%s)\n", cfg.Host, cfg.Name)
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Name,
			cfg.SSLMode,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("open database: %w", err)
		}
		if err := db.AutoMigrate(
			&models.User{},
			&models.Component{},
			// &models.Issue{},
			// &models.Comment{},
		); err != nil {
			return fmt.Errorf("run migrations: %w", err)
		}
		fmt.Println("Migrations applied successfully")

		if autopopulate {
			if err := seedUsers(db); err != nil {
				return fmt.Errorf("seed users: %w", err)
			}
			fmt.Println("Sample users inserted successfully")

			if err := seedComponents(db); err != nil {
				return fmt.Errorf("seed components: %w", err)
			}
			fmt.Println("Sample components inserted successfully")
		}

		return nil
	},
}

func init() {
	migrateCmd.Flags().BoolVar(&autopopulate, "autopopulate", false, "Insert sample data into the users table after migration")
}

func seedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Username: "admin",
			Email:    "admin@bugsbunny.dev",
			Password: "admin123",
			Role:     models.Admin,
			IsActive: true,
		},
		{
			Username: "janedoe",
			Email:    "jane.doe@bugsbunny.dev",
			Password: "jane123",
			Role:     models.Editor,
			IsActive: true,
		},
		{
			Username: "johndoe",
			Email:    "john.doe@bugsbunny.dev",
			Password: "john123",
			Role:     models.Viewer,
			IsActive: true,
		},
	}

	for i := range users {
		result := db.Where("email = ?", users[i].Email).FirstOrCreate(&users[i])
		if result.Error != nil {
			return fmt.Errorf("insert user %q: %w", users[i].Username, result.Error)
		}
	}
	return nil
}

func seedComponents(db *gorm.DB) error {
	components := []models.Component{
		{
			Name:            "General",
			Description:     "All general issues",
			Owner:           "admin",
			Status:          models.ACTIVE,
			SlackChannelID:  nil, // Fixed: do not use a non-existent channel, set to nil
			IsBotEnabled:    false,
			BotKnowledge:    []string{},
			BotInstructions: []string{},
		},
	}

	for i := range components {
		result := db.Where("name = ?", components[i].Name).FirstOrCreate(&components[i])
		if result.Error != nil {
			return fmt.Errorf("insert component %q: %w", components[i].Name, result.Error)
		}
	}
	return nil
}

func loadDBConfig() dbConfig {
	v := viper.New()
	v.SetEnvPrefix("DB")
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		return dbConfig{}
	}

	return dbConfig{
		Host:     v.GetString("DB_HOST"),
		Port:     v.GetString("DB_PORT"),
		User:     v.GetString("DB_USER"),
		Password: v.GetString("DB_PASSWORD"),
		Name:     v.GetString("DB_NAME"),
		SSLMode:  v.GetString("DB_SSLMODE"),
	}
}
