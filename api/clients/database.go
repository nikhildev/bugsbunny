package clients

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func GetDbConfig() (DbConfig, error) {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("DB")
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	if err := v.ReadInConfig(); err != nil {
		return DbConfig{}, fmt.Errorf("failed to read database config: %w", err)
	}
	return DbConfig{
		Host:     v.GetString("DB_HOST"),
		Port:     v.GetString("DB_PORT"),
		User:     v.GetString("DB_USER"),
		Password: v.GetString("DB_PASSWORD"),
		Name:     v.GetString("DB_NAME"),
		SSLMode:  v.GetString("DB_SSL_MODE"),
	}, nil
}

func InitDB(dbConfig DbConfig) (*gorm.DB, error) {
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
