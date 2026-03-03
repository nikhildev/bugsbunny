package updates

import (
	"reflect"
	"strings"
)

// ExtractUpdates builds a map of updates from requestData based on the fields
// present in the model struct. It uses the JSON tags from the struct to match
// against keys in the requestData map.
//
// Example usage:
//
//	updates := updates.ExtractUpdates(requestData, model.Project{})
//
// Parameters:
//   - requestData: map containing the fields from the request body
//   - model: an instance of the struct type (e.g., model.Project{}, model.Issue{})
//
// Returns:
//   - map[string]any: map of field names to values that can be used with GORM's Updates()
func ExtractUpdates[T any](requestData map[string]any, model T) map[string]any {
	result := make(map[string]any)

	// Get the type of the model
	modelType := reflect.TypeOf(model)

	// Handle pointer types
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	// Iterate through all fields of the struct
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)

		// Get the JSON tag for this field
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Handle tags like `json:"name,omitempty"` - extract just the field name
		jsonFieldName := strings.Split(jsonTag, ",")[0]

		// Check if this field exists in the request data
		if val, ok := requestData[jsonFieldName]; ok {
			result[jsonFieldName] = val
		}
	}

	return result
}
