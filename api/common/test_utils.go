package common

import (
	"context"
	"testing"

	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/models"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// SetupTestDB sets up a PostgreSQL testcontainer and initializes the database
// with the necessary schema migrations. Returns the container and a cleanup function.
func SetupTestDB(t *testing.T) (*postgres.PostgresContainer, func()) {
	ctx := context.Background()

	// Create PostgreSQL container
	pgContainer, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2)),
	)
	if err != nil {
		t.Fatalf("Failed to start postgres container: %v", err)
	}

	// Get connection details
	host, err := pgContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}

	port, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get container port: %v", err)
	}

	// Initialize database connection
	dbConfig := clients.DbConfig{
		Host:     host,
		Port:     port.Port(),
		User:     "testuser",
		Password: "testpass",
		Name:     "testdb",
		SSLMode:  "disable",
	}

	db, err := clients.InitDB(dbConfig)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto-migrate schema - Component must come first, then Issue, then Comment (due to foreign keys)
	err = db.AutoMigrate(&models.Component{}, &models.Issue{}, &models.Comment{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// Return cleanup function
	cleanup := func() {
		clients.CloseDbClient()
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}

	return pgContainer, cleanup
}
