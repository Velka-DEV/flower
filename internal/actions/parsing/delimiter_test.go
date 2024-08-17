package parsing

import (
	"reflect"
	"testing"
)

func TestDelimiterParseAction_Execute(t *testing.T) {
	action := &DelimiterParseAction{}
	tests := []struct {
		name           string
		text           string
		startDelimiter string
		endDelimiter   string
		expected       []string
	}{
		{
			name:           "Simple delimiters",
			text:           "Hello [world] and [universe]",
			startDelimiter: "[",
			endDelimiter:   "]",
			expected:       []string{"world", "universe"},
		},
		{
			name:           "Multiple character delimiters",
			text:           "Start{{first}}middle{{second}}end",
			startDelimiter: "{{",
			endDelimiter:   "}}",
			expected:       []string{"first", "second"},
		},
		{
			name:           "No matches",
			text:           "Hello world",
			startDelimiter: "[",
			endDelimiter:   "]",
			expected:       []string{},
		},
		{
			name:           "Nested delimiters",
			text:           "<outer><inner>content</inner></outer>",
			startDelimiter: "<",
			endDelimiter:   ">",
			expected:       []string{"outer", "inner", "/inner", "/outer"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := map[string]interface{}{
				"text":            tt.text,
				"start_delimiter": tt.startDelimiter,
				"end_delimiter":   tt.endDelimiter,
			}

			outputs, err := action.Execute(nil, inputs)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if len(outputs) != 1 || outputs[0].Name != "result" {
				t.Errorf("Unexpected output structure: %v", outputs)
			}

			result, ok := outputs[0].Value.([]string)
			if !ok {
				t.Fatalf("Expected []string, got %T", outputs[0].Value)
			}

			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDelimiterParseAction_Validate(t *testing.T) {
	action := &DelimiterParseAction{}

	validInputs := map[string]interface{}{
		"text":            "Hello [world]",
		"start_delimiter": "[",
		"end_delimiter":   "]",
	}

	if err := action.Validate(validInputs); err != nil {
		t.Errorf("Unexpected error for valid inputs: %v", err)
	}

	invalidInputs := map[string]interface{}{
		"text":            123,
		"start_delimiter": "[",
		"end_delimiter":   "]",
	}

	if err := action.Validate(invalidInputs); err == nil {
		t.Errorf("Expected error for invalid inputs, got nil")
	}
}
