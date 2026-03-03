package cli

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/nikhildev/bugsbunny/api/internal/database"
	"github.com/nikhildev/bugsbunny/api/internal/model"
)

var autopopulate bool
var resetDb bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Apply database migrations to the BugsBunny database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}
		db, err := database.InitDB(cfg.DB)
		if err != nil {
			return fmt.Errorf("init database: %w", err)
		}

		if resetDb {
			if err := db.Migrator().DropTable(&model.User{}, &model.Project{}, &model.Issue{}, &model.Change{}, &model.Comment{}); err != nil {
				return fmt.Errorf("drop tables: %w", err)
			}
			slog.Info("Tables dropped successfully")
		}

		if err := db.AutoMigrate(
			&model.User{},
			&model.Project{},
			&model.Issue{},
			&model.Change{},
			&model.Comment{},
		); err != nil {
			return fmt.Errorf("run migrations: %w", err)
		}

		if err := createDefaultUsers(db); err != nil {
			return fmt.Errorf("create default users: %w", err)
		}

		slog.Info("Migrations applied successfully")

		if autopopulate {
			if err := seedUsers(db); err != nil {
				return fmt.Errorf("seed users: %w", err)
			}
			slog.Info("Sample users inserted successfully")
			if err := seedProjects(db); err != nil {
				return fmt.Errorf("seed projects: %w", err)
			}
			slog.Info("Sample projects inserted successfully")
			if err := seedIssues(db); err != nil {
				return fmt.Errorf("seed issues: %w", err)
			}
			slog.Info("Sample issues inserted successfully")
		}
		return nil
	},
}

func init() {
	migrateCmd.Flags().BoolVar(&autopopulate, "autopopulate", false, "Insert sample data after migration")
	migrateCmd.Flags().BoolVar(&resetDb, "resetdb", false, "Reset the database before running migrations")
}

func createDefaultUsers(db *gorm.DB) error {
	users := []*model.User{
		{Username: "admin", Email: "admin@bugsbunny.dev", Password: "admin123", Role: model.Admin, IsActive: true},
		{Username: "bot", Email: "bot@bugsbunny.dev", Password: "bot123", Role: model.Editor, IsActive: true},
	}
	result := db.Create(users)
	if result.Error != nil {
		slog.Error("Failed to create default users", "error", result.Error)
		return fmt.Errorf("create default users: %w", result.Error)
	}
	slog.Info("Default users created successfully")
	return nil
}

func seedUsers(db *gorm.DB) error {
	users := []model.User{
		{BaseModel: model.BaseModel{ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2c"}, Username: "janedoe", Email: "jane.doe@bugsbunny.dev", Password: "jane123", Role: model.Editor, IsActive: true},
		{BaseModel: model.BaseModel{ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2d"}, Username: "johndoe", Email: "john.doe@bugsbunny.dev", Password: "john123", Role: model.Viewer, IsActive: true},
	}
	for i := range users {
		result := db.Where("id = ?", users[i].ID).FirstOrCreate(&users[i])
		if result.Error != nil {
			slog.Error("Failed to insert user", "user", users[i], "error", result.Error)
			return fmt.Errorf("insert user %q: %w", users[i].Username, result.Error)
		}
	}
	slog.Info("Sample users inserted successfully")
	return nil
}

func seedProjects(db *gorm.DB) error {
	projects := []model.Project{
		{BaseModel: model.BaseModel{ID: "019c48e9-ab2e-7c50-9e03-23f8af4fdd2e"}, Name: "General", Description: "All general issues", Owner: "admin", Status: model.ACTIVE, IsBotEnabled: false, BotKnowledge: []string{}, BotInstructions: []string{}},
	}
	for i := range projects {
		result := db.Where("id = ?", projects[i].ID).FirstOrCreate(&projects[i])
		if result.Error != nil {
			return fmt.Errorf("insert project %q: %w", projects[i].Name, result.Error)
		}
	}
	slog.Info("Sample projects inserted successfully")
	return nil
}

func seedIssues(db *gorm.DB) error {
	issue := model.Issue{
		Title:       "Issue 1",
		Description: "Description 1",
		Status:      model.NEW,
		ReporterId:  uuid.MustParse("019c48e9-ab2e-7c50-9e03-23f8af4fdd2c").String(),
		ProjectID:   uuid.MustParse("019c48e9-ab2e-7c50-9e03-23f8af4fdd2e").String(),
		Type:        model.SUPPORT,
	}
	result := db.Create(&issue)
	if result.Error != nil {
		slog.Error("Failed to insert issue", "error", result.Error)
		return fmt.Errorf("insert issue: %w", result.Error)
	}
	slog.Info("Sample issue inserted successfully")
	return nil
}
