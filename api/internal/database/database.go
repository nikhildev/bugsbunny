package database

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/nikhildev/bugsbunny/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	DB = db
	return db, nil
}

func GetDbClient() (*gorm.DB, error) {
	if DB == nil {
		return nil, errors.New("database not initialized: call InitDB first")
	}
	return DB, nil
}

func CloseDbClient() {
	db, err := DB.DB()
	if err != nil {
		slog.Error("failed to get database client", "error", err)
		return
	}
	if err = db.Close(); err != nil {
		slog.Error("failed to close database client", "error", err)
		return
	}
	slog.Info("database client closed successfully")
}
