package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/models"
)

var autopopulate bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Apply database migrations to the BugsBunny database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := clients.GetDbConfig()
		db, err := clients.InitDB(cfg)
		if err != nil {
			return fmt.Errorf("init database: %w", err)
		}

		if err := db.AutoMigrate(
			&models.User{},
			&models.Component{},
			&models.Issue{},
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

			if err := seedIssues(db); err != nil {
				return fmt.Errorf("seed issues: %w", err)
			}
			fmt.Println("Sample issues inserted successfully")
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
			BaseModel: models.BaseModel{
				ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2b",
			},
			Username: "admin",
			Email:    "admin@bugsbunny.dev",
			Password: "admin123",
			Role:     models.Admin,
			IsActive: true,
		},
		{
			BaseModel: models.BaseModel{
				ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2c",
			},
			Username: "janedoe",
			Email:    "jane.doe@bugsbunny.dev",
			Password: "jane123",
			Role:     models.Editor,
			IsActive: true,
		},
		{
			BaseModel: models.BaseModel{
				ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2d",
			},
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
			BaseModel: models.BaseModel{
				ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2e",
			},
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

func seedIssues(db *gorm.DB) error {
	issues := []models.Issue{
		{
			Title:         "Issue 1",
			Description:   "Description 1",
			Status:        models.NEW,
			Reporter:      "019c48e9-ab2e-7c50-9e03-23f8af4fdd2c",
			ComponentID:   "019c48e9-ab2e-7c50-9e03-23f8af4fdd2e",
			Type:          models.SUPPORT,
			Attachments:   []string{},
			Priority:      models.LOW_PRIORITY,
			Severity:      models.LOW_SEVERITY,
			Collaborators: []string{},
			CC:            []string{},
		},
	}

	for i := range issues {
		result := db.Where("title = ?", issues[i].Title).FirstOrCreate(&issues[i])
		if result.Error != nil {
			return fmt.Errorf("insert issue %q: %w", issues[i].Title, result.Error)
		}
	}
	return nil
}
