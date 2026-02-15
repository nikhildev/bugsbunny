# Generic Update Extractor

The `ExtractUpdates` function uses Go generics to provide a type-safe way to extract update fields from request data for any struct type.

## Usage

Instead of manually checking each field:

```go
// Old approach - repetitive and error-prone
updates := make(map[string]any)
if val, ok := requestData["name"]; ok {
    updates["name"] = val
}
if val, ok := requestData["description"]; ok {
    updates["description"] = val
}
// ... many more lines
```

Use the generic function:

```go
// New approach - works with any model
updates := common.ExtractUpdates(requestData, models.Component{})
```

## How It Works

The function uses Go generics and reflection to:
1. Accept any struct type as a type parameter
2. Read the struct's JSON tags using reflection
3. Match them against keys in the request data map
4. Build a map of only the fields present in the request

This provides compile-time type safety and is perfect for GORM's `Updates()` method, which needs a map to perform partial updates.

## Examples

### With Component Model

```go
requestData := map[string]any{
    "name": "New Name",
    "owner": "john@example.com",
}

updates := common.ExtractUpdates(requestData, models.Component{})
db.Model(&models.Component{}).Where("id = ?", id).Updates(updates)
```

### With Issue Model

```go
requestData := map[string]any{
    "title": "Bug fix",
    "status": "in_progress",
    "assignee": "jane@example.com",
}

updates := common.ExtractUpdates(requestData, models.Issue{})
db.Model(&models.Issue{}).Where("id = ?", id).Updates(updates)
```

### With Any Custom Model

As long as your struct has JSON tags, it will work:

```go
type CustomModel struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}

updates := common.ExtractUpdates(requestData, CustomModel{})
```

## Benefits

- **DRY**: No repetitive field checking code
- **Type-safe**: Uses Go generics for compile-time type safety
- **Flexible**: Works with any struct that has JSON tags
- **Maintainable**: Adding new fields to a model automatically works
- **Testable**: Easy to unit test (see `updates_test.go`)
- **GORM-friendly**: Output works directly with GORM's `Updates()` method
