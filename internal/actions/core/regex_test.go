package core

import (
	"testing"
)

func TestRegexAction_Execute(t *testing.T) {
	action := &RegexAction{}

	inputs := map[string]interface{}{
		"regex": `\d+`,
		"text":  "The number is 42",
	}

	outputs, err := action.Execute(nil, inputs)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(outputs) != 2 {
		t.Fatalf("Expected 2 outputs, got %d", len(outputs))
	}

	matches, ok := outputs[0].Value.([]string)
	if !ok || len(matches) != 1 || matches[0] != "42" {
		t.Errorf("Unexpected matches: %v", outputs[0].Value)
	}
}

func TestRegexAction_Validate(t *testing.T) {
	action := &RegexAction{}

	tests := []struct {
		name    string
		inputs  map[string]interface{}
		wantErr bool
	}{
		{
			name: "Valid inputs",
			inputs: map[string]interface{}{
				"regex": `\d+`,
				"text":  "The number is 42",
			},
			wantErr: false,
		},
		{
			name: "Missing regex",
			inputs: map[string]interface{}{
				"text": "The number is 42",
			},
			wantErr: true,
		},
		{
			name: "Invalid regex type",
			inputs: map[string]interface{}{
				"regex": 123,
				"text":  "The number is 42",
			},
			wantErr: true,
		},
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
