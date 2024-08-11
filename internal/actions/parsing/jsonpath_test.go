package parsing

import (
	"reflect"
	"testing"
)

func TestJSONPathAction_Execute(t *testing.T) {
	action := &JSONPathAction{}
	tests := []struct {
		name     string
		json     string
		path     string
		expected interface{}
	}{
		{
			name:     "Simple path",
			json:     `{"name": "John", "age": 30, "city": "New York"}`,
			path:     "$.name",
			expected: "John",
		},
		{
			name:     "Array index",
			json:     `{"people": [{"name": "John"}, {"name": "Jane"}]}`,
			path:     "$.people[1].name",
			expected: "Jane",
		},
		{
			name:     "Multiple results",
			json:     `{"people": [{"name": "John"}, {"name": "Jane"}]}`,
			path:     "$.people[*].name",
			expected: []interface{}{"John", "Jane"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := map[string]interface{}{
				"json": tt.json,
				"path": tt.path,
			}

			outputs, err := action.Execute(nil, inputs)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(outputs) != 1 || outputs[0].Name != "result" {
				t.Errorf("Unexpected output structure: %v", outputs)
			}

			if !reflect.DeepEqual(outputs[0].Value, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, outputs[0].Value)
			}
		})
	}
}

func TestJSONPathAction_Validate(t *testing.T) {
	action := &JSONPathAction{}

	validInputs := map[string]interface{}{
		"json": `{"name": "John", "age": 30}`,
		"path": "$.name",
	}

	if err := action.Validate(validInputs); err != nil {
		t.Errorf("Unexpected error for valid inputs: %v", err)
	}

	invalidInputs := map[string]interface{}{
		"json": 123,
		"path": "$.name",
	}

	if err := action.Validate(invalidInputs); err == nil {
		t.Errorf("Expected error for invalid inputs, got nil")
	}
}
