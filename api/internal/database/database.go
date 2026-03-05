package database

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbConfig config.DbConfig) (*gorm.DB, error) {
	slog.Info("initializing database...")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User,
		dbConfig.Password, dbConfig.Name, dbConfig.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database: " + err.Error())
	}
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to get database client", "error", err)
		return
	}
	if err = sqlDB.Close(); err != nil {
		slog.Error("failed to close database client", "error", err)
		return
	}
	slog.Info("database client closed successfully")
}
