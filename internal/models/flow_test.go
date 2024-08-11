package models

import (
	"testing"
)

func TestFromYaml(t *testing.T) {
	yamlData := []byte(`
name: Test Flow
description: A test flow
version: 1.0.0
author: Tester
flow_version: 0.0.1
steps:
  - id: step1
    name: Step 1
    action: test/action
    inputs:
      input1: value1
`)

	flow, err := FromYaml(yamlData)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if flow.Name != "Test Flow" {
		t.Errorf("Expected flow name 'Test Flow', got '%s'", flow.Name)
	}

	if len(flow.Steps) != 1 {
		t.Fatalf("Expected 1 step, got %d", len(flow.Steps))
	}

	if flow.Steps[0].ID != "step1" {
		t.Errorf("Expected step ID 'step1', got '%s'", flow.Steps[0].ID)
	}
}

func TestFlow_ToYaml(t *testing.T) {
	flow := &Flow{
		Name:        "Test Flow",
		Description: "A test flow",
		Version:     "1.0.0",
		Author:      "Tester",
		FlowVersion: "0.0.1",
		Steps: []Step{
			{
				ID:     "step1",
				Name:   "Step 1",
				Action: "test/action",
				Inputs: map[string]interface{}{
					"input1": "value1",
				},
			},
		},
	}

	yamlData, err := flow.ToYaml()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Convert back to Flow to check if all data is preserved
	newFlow, err := FromYaml(yamlData)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if newFlow.Name != flow.Name {
		t.Errorf("Expected flow name '%s', got '%s'", flow.Name, newFlow.Name)
	}

	if len(newFlow.Steps) != len(flow.Steps) {
		t.Fatalf("Expected %d steps, got %d", len(flow.Steps), len(newFlow.Steps))
	}

	if newFlow.Steps[0].ID != flow.Steps[0].ID {
		t.Errorf("Expected step ID '%s', got '%s'", flow.Steps[0].ID, newFlow.Steps[0].ID)
	}
}
