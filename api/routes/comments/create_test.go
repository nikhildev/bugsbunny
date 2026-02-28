package comments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/nikhildev/bugsbunny/api/clients"
	"github.com/nikhildev/bugsbunny/api/common"
	"github.com/nikhildev/bugsbunny/api/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommentHandler_Success(t *testing.T) {
	// Setup
	_, cleanup := common.SetupTestDB(t)
	defer cleanup()

	// Create test data: Component and Issue (required for foreign keys)
	db, _ := clients.GetDbClient()

	// Create Component first - BeforeCreate checks if ID is empty
	component := models.Component{
		Name:        "Test Component",
		Description: "Test Description",
		Owner:       "test-owner",
		Status:      models.ACTIVE,
	}
	result := db.Create(&component)
	if result.Error != nil {
		t.Fatalf("Failed to create component: %v", result.Error)
	}

	// Create Issue - BeforeCreate will generate ID automatically
	// Create users for assignee and reporter
	assignee := models.User{
		Username: "test-assignee",
		Email:    "assignee@test.com",
		Password: "password",
		Role:     models.Editor,
		IsActive: true,
	}
	db.Create(&assignee)

	reporter := models.User{
		Username: "test-reporter",
		Email:    "reporter@test.com",
		Password: "password",
		Role:     models.Editor,
		IsActive: true,
	}
	db.Create(&reporter)

	issue := models.Issue{
		Title:       "Test Issue",
		Description: "Test Description",
		Type:        models.BUG,
		Status:      models.NEW,
		AssigneeId:  &assignee.ID,
		ReporterId:  reporter.ID,
		ComponentID: component.ID,
		Priority:    models.LOW_PRIORITY,
		Severity:    models.LOW_SEVERITY,
	}
	issueResult := db.Create(&issue)
	if issueResult.Error != nil {
		t.Fatalf("Failed to create issue: %v", issueResult.Error)
	}

	// Use the actual issue ID that was generated
	issueID := issue.ID

	// Prepare request body
	authorUUID, _ := uuid.NewV7()
	comment := models.Comment{
		Content:     "This is a test comment",
		Attachments: []string{"file1.png", "file2.jpg"},
	}
	body, _ := json.Marshal(comment)

	// Create request with path parameter
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader(body))
	req.SetPathValue("id", issueID)
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute handler
	CreateCommentHandler(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code, "Expected status code 201")

	// Verify comment was created in database
	var createdComment models.Comment
	queryResult := db.Where("issue_id = ?", issueID).First(&createdComment)
	assert.NoError(t, queryResult.Error, "Comment should be created in database")
	assert.Equal(t, "This is a test comment", createdComment.Content)
	assert.Equal(t, issueID, createdComment.IssueID)
	assert.Equal(t, authorUUID, createdComment.Author, "Author should match the X-User-UUID header")
	assert.Equal(t, 2, len(createdComment.Attachments))
}

func TestCreateCommentHandler_MissingIssueID(t *testing.T) {
	// Setup
	_, cleanup := common.SetupTestDB(t)
	defer cleanup()

	// Prepare request body
	authorUUID, _ := uuid.NewV7()
	comment := models.Comment{
		Content: "This is a test comment",
	}
	body, _ := json.Marshal(comment)

	// Create request without path parameter
	req := httptest.NewRequest(http.MethodPost, "/issues//comments", bytes.NewReader(body))
	// Don't set path value - simulating missing issue ID
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute handler
	CreateCommentHandler(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for missing issue ID")
}

func TestCreateCommentHandler_InvalidJSON(t *testing.T) {
	// Setup
	_, cleanup := common.SetupTestDB(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	issueIDStr := issueID.String()
	authorUUID, _ := uuid.NewV7()

	// Create request with invalid JSON
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueIDStr), bytes.NewReader([]byte("invalid json")))
	req.SetPathValue("id", issueIDStr)
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute handler
	CreateCommentHandler(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for invalid JSON")
}

func TestCreateCommentHandler_InvalidUserUUID(t *testing.T) {
	// Setup
	_, cleanup := common.SetupTestDB(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	issueIDStr := issueID.String()

	// Prepare request body
	comment := models.Comment{
		Content: "This is a test comment",
	}
	body, _ := json.Marshal(comment)

	// Create request with invalid UUID header
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueIDStr), bytes.NewReader(body))
	req.SetPathValue("id", issueIDStr)
	req.Header.Set("X-User-UUID", "invalid-uuid")
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute handler
	CreateCommentHandler(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for invalid UUID")
}

func TestCreateCommentHandler_MissingUserUUID(t *testing.T) {
	// Setup
	_, cleanup := common.SetupTestDB(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	issueIDStr := issueID.String()

	// Prepare request body
	comment := models.Comment{
		Content: "This is a test comment",
	}
	body, _ := json.Marshal(comment)

	// Create request without X-User-UUID header
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueIDStr), bytes.NewReader(body))
	req.SetPathValue("id", issueIDStr)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute handler
	CreateCommentHandler(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for missing UUID header")
}

func TestCreateCommentHandler_EmptyContent(t *testing.T) {
	// Setup
	_, cleanup := common.SetupTestDB(t)
	defer cleanup()

	// Create test data: Component and Issue (required for foreign keys)
	db, _ := clients.GetDbClient()

	// Create Component first - BeforeCreate checks if ID is empty
	component := models.Component{
		Name:        "Test Component",
		Description: "Test Description",
		Owner:       "test-owner",
		Status:      models.ACTIVE,
	}
	result := db.Create(&component)
	if result.Error != nil {
		t.Fatalf("Failed to create component: %v", result.Error)
	}

	// Create users for assignee and reporter
	assignee := models.User{
		Username: "test-assignee",
		Email:    "assignee@test.com",
		Password: "password",
		Role:     models.Editor,
		IsActive: true,
	}
	db.Create(&assignee)

	reporter := models.User{
		Username: "test-reporter",
		Email:    "reporter@test.com",
		Password: "password",
		Role:     models.Editor,
		IsActive: true,
	}
	db.Create(&reporter)

	// Create Issue - BeforeCreate will generate ID automatically
	issue := models.Issue{
		Title:       "Test Issue",
		Description: "Test Description",
		Type:        models.BUG,
		Status:      models.NEW,
		AssigneeId:  &assignee.ID,
		ReporterId:  reporter.ID,
		ComponentID: component.ID,
		Priority:    models.LOW_PRIORITY,
		Severity:    models.LOW_SEVERITY,
	}
	issueResult := db.Create(&issue)
	if issueResult.Error != nil {
		t.Fatalf("Failed to create issue: %v", issueResult.Error)
	}

	// Use the actual issue ID that was generated
	issueID := issue.ID

	// Prepare request body with empty content
	authorUUID, _ := uuid.NewV7()
	comment := models.Comment{
		Content: "",
	}
	body, _ := json.Marshal(comment)

	// Create request
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader(body))
	req.SetPathValue("id", issueID)
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute handler
	CreateCommentHandler(w, req)

	// Note: This test documents current behavior - the handler accepts empty content
	// If validation is added later, this test should be updated
	assert.Equal(t, http.StatusCreated, w.Code, "Current implementation accepts empty content")
}
