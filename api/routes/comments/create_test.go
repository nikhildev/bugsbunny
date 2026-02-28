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

func newTestHandler(t *testing.T) (*Handler, func()) {
	t.Helper()
	_, cleanup := common.SetupTestDB(t)
	db, err := clients.GetDbClient()
	if err != nil {
		t.Fatalf("failed to get db client: %v", err)
	}
	return &Handler{DB: db}, cleanup
}

func TestCreateCommentHandler_Success(t *testing.T) {
	h, cleanup := newTestHandler(t)
	defer cleanup()

	db, _ := clients.GetDbClient()

	component := models.Component{
		Name:        "Test Component",
		Description: "Test Description",
		Owner:       "test-owner",
		Status:      models.ACTIVE,
	}
	if result := db.Create(&component); result.Error != nil {
		t.Fatalf("Failed to create component: %v", result.Error)
	}

	issue := models.Issue{
		Title:       "Test Issue",
		Description: "Test Description",
		Type:        models.BUG,
		Status:      models.NEW,
		ReporterId:  "019c48e9-ab2e-7c50-9e03-23f8af4fdd2c",
		ComponentID: component.ID,
		Priority:    models.LOW_PRIORITY,
		Severity:    models.LOW_SEVERITY,
	}
	if result := db.Create(&issue); result.Error != nil {
		t.Fatalf("Failed to create issue: %v", result.Error)
	}

	issueID := issue.ID
	authorUUID, _ := uuid.NewV7()
	comment := models.Comment{
		Content:     "This is a test comment",
		Attachments: []string{"file1.png", "file2.jpg"},
	}
	body, _ := json.Marshal(comment)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader(body))
	req.SetPathValue("id", issueID)
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateComment(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status code 201")

	var createdComment models.Comment
	queryResult := db.Where("issue_id = ?", issueID).First(&createdComment)
	assert.NoError(t, queryResult.Error, "Comment should be created in database")
	assert.Equal(t, "This is a test comment", createdComment.Content)
	assert.Equal(t, issueID, createdComment.IssueID)
	assert.Equal(t, authorUUID, createdComment.Author)
	assert.Equal(t, 2, len(createdComment.Attachments))
}

func TestCreateCommentHandler_MissingIssueID(t *testing.T) {
	h, cleanup := newTestHandler(t)
	defer cleanup()

	authorUUID, _ := uuid.NewV7()
	comment := models.Comment{Content: "This is a test comment"}
	body, _ := json.Marshal(comment)

	req := httptest.NewRequest(http.MethodPost, "/issues//comments", bytes.NewReader(body))
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateComment(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for missing issue ID")
}

func TestCreateCommentHandler_InvalidJSON(t *testing.T) {
	h, cleanup := newTestHandler(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	authorUUID, _ := uuid.NewV7()

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader([]byte("invalid json")))
	req.SetPathValue("id", issueID.String())
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateComment(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for invalid JSON")
}

func TestCreateCommentHandler_InvalidUserUUID(t *testing.T) {
	h, cleanup := newTestHandler(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	comment := models.Comment{Content: "This is a test comment"}
	body, _ := json.Marshal(comment)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader(body))
	req.SetPathValue("id", issueID.String())
	req.Header.Set("X-User-UUID", "invalid-uuid")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateComment(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for invalid UUID")
}

func TestCreateCommentHandler_MissingUserUUID(t *testing.T) {
	h, cleanup := newTestHandler(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	comment := models.Comment{Content: "This is a test comment"}
	body, _ := json.Marshal(comment)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader(body))
	req.SetPathValue("id", issueID.String())
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateComment(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for missing UUID header")
}

func TestCreateCommentHandler_EmptyContent(t *testing.T) {
	h, cleanup := newTestHandler(t)
	defer cleanup()

	issueID, _ := uuid.NewV7()
	authorUUID, _ := uuid.NewV7()
	comment := models.Comment{Content: ""}
	body, _ := json.Marshal(comment)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/issues/%s/comments", issueID), bytes.NewReader(body))
	req.SetPathValue("id", issueID.String())
	req.Header.Set("X-User-UUID", authorUUID.String())
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateComment(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 for empty content")
}
