package core

import (
	"testing"
)

func TestPrintAction_Execute(t *testing.T) {
	action := &PrintAction{}

	inputs := map[string]interface{}{
		"message": "Hello, World!",
	}

	outputs, err := action.Execute(nil, inputs)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(outputs) != 0 {
		t.Errorf("Expected 0 outputs, got %d", len(outputs))
	}
}

func TestPrintAction_Validate(t *testing.T) {
	action := &PrintAction{}

	tests := []struct {
		name    string
		inputs  map[string]interface{}
		wantErr bool
	}{
		{
			name: "Valid input",
			inputs: map[string]interface{}{
				"message": "Hello, World!",
			},
			wantErr: false,
		},
		{
			name:    "Missing message",
			inputs:  map[string]interface{}{},
			wantErr: true,
		},
		{
			name: "Invalid message type",
			inputs: map[string]interface{}{
				"message": 123,
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
