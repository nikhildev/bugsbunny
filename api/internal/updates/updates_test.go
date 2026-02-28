package updates

import (
	"testing"

	"github.com/nikhildev/bugsbunny/api/internal/model"
)

func TestExtractUpdates_Component(t *testing.T) {
	requestData := map[string]any{
		"name":              "Updated Component",
		"description":       "New description",
		"owner":             "john@example.com",
		"is_bot_enabled":    true,
		"non_existent_field": "should be ignored",
	}

	updates := ExtractUpdates(requestData, model.Component{})

	if len(updates) != 4 {
		t.Errorf("Expected 4 updates, got %d", len(updates))
	}

	if updates["name"] != "Updated Component" {
		t.Errorf("Expected name to be 'Updated Component', got %v", updates["name"])
	}

	if updates["description"] != "New description" {
		t.Errorf("Expected description to be 'New description', got %v", updates["description"])
	}

	if updates["owner"] != "john@example.com" {
		t.Errorf("Expected owner to be 'john@example.com', got %v", updates["owner"])
	}

	if updates["is_bot_enabled"] != true {
		t.Errorf("Expected is_bot_enabled to be true, got %v", updates["is_bot_enabled"])
	}

	if _, exists := updates["non_existent_field"]; exists {
		t.Errorf("Non-existent field should not be in updates")
	}
}

func TestExtractUpdates_Issue(t *testing.T) {
	requestData := map[string]any{
		"title":       "Bug in login",
		"description": "Users cannot login",
		"status":      "in_progress",
		"assignee":    "jane@example.com",
		"priority":    "high",
		"invalid":     "should be ignored",
	}

	updates := ExtractUpdates(requestData, model.Issue{})

	if len(updates) != 5 {
		t.Errorf("Expected 5 updates, got %d", len(updates))
	}

	if updates["title"] != "Bug in login" {
		t.Errorf("Expected title to be 'Bug in login', got %v", updates["title"])
	}

	if updates["status"] != "in_progress" {
		t.Errorf("Expected status to be 'in_progress', got %v", updates["status"])
	}

	if _, exists := updates["invalid"]; exists {
		t.Errorf("Invalid field should not be in updates")
	}
}

func TestExtractUpdates_EmptyRequestData(t *testing.T) {
	requestData := map[string]any{}

	updates := ExtractUpdates(requestData, model.Component{})

	if len(updates) != 0 {
		t.Errorf("Expected 0 updates for empty request data, got %d", len(updates))
	}
}

func TestExtractUpdates_OnlyNonMatchingFields(t *testing.T) {
	requestData := map[string]any{
		"invalid_field_1": "value1",
		"invalid_field_2": "value2",
	}

	updates := ExtractUpdates(requestData, model.Component{})

	if len(updates) != 0 {
		t.Errorf("Expected 0 updates when no fields match, got %d", len(updates))
	}
}
