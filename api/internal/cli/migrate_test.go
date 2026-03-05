package cli

import (
	"testing"

	"github.com/nikhildev/bugsbunny/api/internal/database"
	"github.com/nikhildev/bugsbunny/api/internal/model"
	"github.com/nikhildev/bugsbunny/api/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupMigrateTestDB(t *testing.T) (*gorm.DB, func()) {
	t.Helper()
	_, cleanup := testutil.SetupTestDB(t)
	db, _ := database.GetDbClient()
	err := db.AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Issue{},
		&model.Change{},
		&model.Comment{},
	)
	require.NoError(t, err)
	return db, cleanup
}

func TestCreateDefaultUsers(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	err := createDefaultUsers(db)
	require.NoError(t, err)

	var users []model.User
	db.Find(&users)
	assert.Len(t, users, 2)

	var admin model.User
	db.Where("username = ?", "admin").First(&admin)
	assert.Equal(t, "admin@bugsbunny.dev", admin.Email)
	assert.Equal(t, model.Admin, admin.Role)
	assert.True(t, admin.IsActive)

	var bot model.User
	db.Where("username = ?", "bot").First(&bot)
	assert.Equal(t, "bot@bugsbunny.dev", bot.Email)
	assert.Equal(t, model.Editor, bot.Role)
	assert.True(t, bot.IsActive)
}

func TestCreateDefaultUsers_Duplicate(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	err := createDefaultUsers(db)
	require.NoError(t, err)

	err = createDefaultUsers(db)
	assert.Error(t, err, "should fail on duplicate default users")
}

func TestSeedFromJSON_Users(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	err := seedFromJSON[model.User](db, "seeddata/users.json", true)
	require.NoError(t, err)

	var users []model.User
	db.Find(&users)
	assert.Len(t, users, 2)

	var jane model.User
	db.Where("username = ?", "janedoe").First(&jane)
	assert.Equal(t, "019c48e9-ab2e-7c50-9e03-23f8af4fdd2c", jane.ID)
	assert.Equal(t, "jane.doe@bugsbunny.dev", jane.Email)
	assert.Equal(t, model.Editor, jane.Role)
	assert.True(t, jane.IsActive)

	var john model.User
	db.Where("username = ?", "johndoe").First(&john)
	assert.Equal(t, "019c48e9-ab2e-7c50-9e03-23f8af4fdd2d", john.ID)
	assert.Equal(t, "john.doe@bugsbunny.dev", john.Email)
	assert.Equal(t, model.Viewer, john.Role)
	assert.True(t, john.IsActive)
}

func TestSeedFromJSON_Users_SkipConflicts(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	err := seedFromJSON[model.User](db, "seeddata/users.json", true)
	require.NoError(t, err)

	// Seed again — should not error due to ON CONFLICT DO NOTHING
	err = seedFromJSON[model.User](db, "seeddata/users.json", true)
	require.NoError(t, err)

	var count int64
	db.Model(&model.User{}).Count(&count)
	assert.Equal(t, int64(2), count, "duplicate seed should not create extra rows")
}

func TestSeedFromJSON_Projects(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	err := seedFromJSON[model.Project](db, "seeddata/projects.json", true)
	require.NoError(t, err)

	var projects []model.Project
	db.Find(&projects)
	require.Len(t, projects, 1)
	assert.Equal(t, "019c48e9-ab2e-7c50-9e03-23f8af4fdd2e", projects[0].ID)
	assert.Equal(t, "General", projects[0].Name)
	assert.Equal(t, "All general issues", projects[0].Description)
	assert.Equal(t, "admin", projects[0].Owner)
	assert.Equal(t, model.ProjectStatusActive, projects[0].Status)
	assert.False(t, projects[0].IsBotEnabled)
}

func TestSeedFromJSON_Issues(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	// Seed users and projects first (foreign key dependencies)
	require.NoError(t, seedFromJSON[model.User](db, "seeddata/users.json", true))
	require.NoError(t, seedFromJSON[model.Project](db, "seeddata/projects.json", true))

	err := seedFromJSON[model.Issue](db, "seeddata/issues.json", false)
	require.NoError(t, err)

	var issues []model.Issue
	db.Find(&issues)
	require.Len(t, issues, 1)
	assert.Equal(t, "Issue 1", issues[0].Title)
	assert.Equal(t, "Description 1", issues[0].Description)
	assert.Equal(t, model.IssueStatusNew, issues[0].Status)
	assert.Equal(t, model.IssueTypeSupport, issues[0].Type)
	assert.Equal(t, "019c48e9-ab2e-7c50-9e03-23f8af4fdd2c", issues[0].ReporterId)
	assert.Equal(t, "019c48e9-ab2e-7c50-9e03-23f8af4fdd2e", issues[0].ProjectID)
}

func TestSeedFromJSON_Issues_NoDuplicateProtection(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	require.NoError(t, seedFromJSON[model.User](db, "seeddata/users.json", true))
	require.NoError(t, seedFromJSON[model.Project](db, "seeddata/projects.json", true))

	require.NoError(t, seedFromJSON[model.Issue](db, "seeddata/issues.json", false))
	require.NoError(t, seedFromJSON[model.Issue](db, "seeddata/issues.json", false))

	var count int64
	db.Model(&model.Issue{}).Count(&count)
	assert.Equal(t, int64(2), count, "issues without skipConflicts should create duplicates")
}

func TestSeedFromJSON_InvalidPath(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	err := seedFromJSON[model.User](db, "seeddata/nonexistent.json", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read embedded file")
}

func TestSeedFromJSON_FullSeedSequence(t *testing.T) {
	db, cleanup := setupMigrateTestDB(t)
	defer cleanup()

	// Simulate the full autopopulate sequence from the migrate command
	require.NoError(t, createDefaultUsers(db))
	require.NoError(t, seedFromJSON[model.User](db, "seeddata/users.json", true))
	require.NoError(t, seedFromJSON[model.Project](db, "seeddata/projects.json", true))
	require.NoError(t, seedFromJSON[model.Issue](db, "seeddata/issues.json", false))

	var userCount int64
	db.Model(&model.User{}).Count(&userCount)
	assert.Equal(t, int64(4), userCount, "2 default + 2 seed users")

	var projectCount int64
	db.Model(&model.Project{}).Count(&projectCount)
	assert.Equal(t, int64(1), projectCount)

	var issueCount int64
	db.Model(&model.Issue{}).Count(&issueCount)
	assert.Equal(t, int64(1), issueCount)
}
