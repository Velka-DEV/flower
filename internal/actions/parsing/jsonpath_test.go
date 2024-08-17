package parsing_test

import (
	"flower/internal/actions/parsing"
	"reflect"
	"testing"
)

func TestJSONPathAction_GetIdentifier(t *testing.T) {
	action := parsing.JSONPathAction{}
	if action.GetIdentifier() != "parsing/jsonpath" {
		t.Errorf("Expected identifier 'parsing/jsonpath', got %s", action.GetIdentifier())
	}
}

func TestJSONPathAction_GetInputSchema(t *testing.T) {
	action := parsing.JSONPathAction{}
	schema := action.GetInputSchema()
	if len(schema) != 2 {
		t.Errorf("Expected 2 inputs, got %d", len(schema))
	}
	if schema[0].Name != "json" || schema[1].Name != "path" {
		t.Errorf("Unexpected input names: %v", schema)
	}
}

func TestJSONPathAction_GetOutputSchema(t *testing.T) {
	action := parsing.JSONPathAction{}
	schema := action.GetOutputSchema()
	if len(schema) != 1 {
		t.Errorf("Expected 1 output, got %d", len(schema))
	}
	if schema[0].Name != "result" {
		t.Errorf("Unexpected output name: %s", schema[0].Name)
	}
}

func TestJSONPathAction_Validate(t *testing.T) {
	action := parsing.JSONPathAction{}
	tests := []struct {
		name    string
		inputs  map[string]interface{}
		wantErr bool
	}{
		{"Valid inputs", map[string]interface{}{"json": "{}", "path": "$"}, false},
		{"Missing json", map[string]interface{}{"path": "$"}, true},
		{"Missing path", map[string]interface{}{"json": "{}"}, true},
		{"Non-string json", map[string]interface{}{"json": 123, "path": "$"}, true},
		{"Non-string path", map[string]interface{}{"json": "{}", "path": 123}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := action.Validate(tt.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONPathAction_Execute(t *testing.T) {
	action := parsing.JSONPathAction{}
	testContext := &testContext{}

	tests := []struct {
		name    string
		inputs  map[string]interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:   "Simple JSON object",
			inputs: map[string]interface{}{"json": `{"name": "John", "age": 30}`, "path": "$.name"},
			want:   "John",
		},
		{
			name:   "Nested JSON object",
			inputs: map[string]interface{}{"json": `{"person": {"name": "John", "age": 30}}`, "path": "$.person.age"},
			want:   int64(30),
		},
		{
			name:   "JSON array",
			inputs: map[string]interface{}{"json": `[1, 2, 3, 4, 5]`, "path": "$[2]"},
			want:   int64(3),
		},
		{
			name:   "Wildcard",
			inputs: map[string]interface{}{"json": `{"people": [{"name": "John"}, {"name": "Jane"}]}`, "path": "$.people[*].name"},
			want:   []interface{}{"John", "Jane"},
		},
		{
			name:   "Filter expression",
			inputs: map[string]interface{}{"json": `[{"name": "John", "age": 30}, {"name": "Jane", "age": 25}]`, "path": "$[?(@.age > 28)]"},
			want:   map[string]interface{}{"name": "John", "age": int64(30)},
		},
		{
			name:    "Invalid JSON",
			inputs:  map[string]interface{}{"json": `{"invalid": "json"`, "path": "$"},
			wantErr: true,
		},
		{
			name:    "Invalid JSONPath",
			inputs:  map[string]interface{}{"json": `{}`, "path": "$[invalid]"},
			wantErr: true,
		},
		{
			name:   "Non-existent path",
			inputs: map[string]interface{}{"json": `{"name": "John"}`, "path": "$.age"},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputs, err := action.Execute(testContext, tt.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if len(outputs) != 1 {
				t.Errorf("Expected 1 output, got %d", len(outputs))
				return
			}
			if !reflect.DeepEqual(outputs[0].Value, tt.want) {
				t.Errorf("Execute() got = %v (%T), want %v (%T)", outputs[0].Value, outputs[0].Value, tt.want, tt.want)
			}
		})
	}
}

// Mock Context for testing
type testContext struct{}

func (c *testContext) SetOutput(stepID, outputName string, value interface{}) {}
