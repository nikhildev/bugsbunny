package cli

import (
	"embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"github.com/nikhildev/bugsbunny/api/internal/database"
	"github.com/nikhildev/bugsbunny/api/internal/model"
)

//go:embed seeddata/users.json seeddata/projects.json seeddata/issues.json
var seedFS embed.FS

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
			if err := seedFromJSON[model.User](db, "seeddata/users.json", true); err != nil {
				return fmt.Errorf("seed users: %w", err)
			}
			slog.Info("Sample users inserted successfully")
			if err := seedFromJSON[model.Project](db, "seeddata/projects.json", true); err != nil {
				return fmt.Errorf("seed projects: %w", err)
			}
			slog.Info("Sample projects inserted successfully")
			if err := seedFromJSON[model.Issue](db, "seeddata/issues.json", false); err != nil {
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

// seedFromJSON reads a JSON file from the embedded filesystem and batch-inserts
// the records into the database. When skipConflicts is true, existing rows
// (matched by primary key) are silently skipped via ON CONFLICT DO NOTHING.
func seedFromJSON[T any](db *gorm.DB, path string, skipConflicts bool) error {
	data, err := seedFS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read embedded file %s: %w", path, err)
	}

	var records []T
	if err := json.Unmarshal(data, &records); err != nil {
		return fmt.Errorf("unmarshal %s: %w", path, err)
	}

	tx := db
	if skipConflicts {
		tx = tx.Clauses(clause.OnConflict{DoNothing: true})
	}
	if result := tx.Create(&records); result.Error != nil {
		return fmt.Errorf("insert records from %s: %w", path, result.Error)
	}
	return nil
}
