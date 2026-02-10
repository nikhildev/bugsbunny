package clients

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the global database connection instance, set by InitDB.
var DB *gorm.DB

// DbConfig contains the parameters needed to connect to a PostgreSQL database.
type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// GetDbConfig reads database connection settings from environment variables
// prefixed with "DB_" (e.g. DB_HOST, DB_PORT) using Viper and returns a
// populated DbConfig. If the configuration cannot be read, it prints an error
// and returns a zero-value DbConfig.
func GetDbConfig() DbConfig {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("DB")
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()

	if err != nil {
		fmt.Println("failed to read database config: " + err.Error())
		return DbConfig{}
	}

	return DbConfig{
		Host:     v.GetString("DB_HOST"),
		Port:     v.GetString("DB_PORT"),
		User:     v.GetString("DB_USER"),
		Password: v.GetString("DB_PASSWORD"),
		Name:     v.GetString("DB_NAME"),
		SSLMode:  v.GetString("DB_SSL_MODE"),
	}
}

// InitDB opens a PostgreSQL connection using the provided DbConfig, stores it
// in the package-level DB variable, and returns the *gorm.DB handle. Returns an
// error if the connection cannot be established.
func InitDB(dbConfig DbConfig) (*gorm.DB, error) {
	fmt.Println("initializing database...")
	fmt.Println("database config: ", dbConfig)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.SSLMode,
	)
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database: " + err.Error())
	}
	DB = db
	return db, nil
}

// GetDbClient returns the shared *gorm.DB instance initialized by InitDB.
// Returns an error if InitDB has not been called yet.
func GetDbClient() (*gorm.DB, error) {
	if DB == nil {
		return nil, errors.New("database not initialized: call InitDB first")
	}
	return DB, nil
}

// CloseDbClient gracefully closes the underlying sql.DB connection held by the
// package-level DB instance. Errors during retrieval or closing are logged to
// stdout.
func CloseDbClient() {
	db, err := DB.DB()
	if err != nil {
		fmt.Println("failed to get database client: " + err.Error())
		return
	}
	err = db.Close()
	if err != nil {
		fmt.Println("failed to close database client: " + err.Error())
		return
	}
	fmt.Println("database client closed successfully")
}
