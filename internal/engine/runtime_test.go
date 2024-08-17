package engine

import (
	"flower/internal/actions"
	"flower/internal/models"
	"testing"
)

func TestRuntime_Run(t *testing.T) {
	// Create a mock action
	mockAction := &MockAction{
		identifier: "mock/action",
		executeFunc: func(ctx models.Context, inputs map[string]interface{}) ([]models.Output, error) {
			return []models.Output{{Name: "result", Value: "success"}}, nil
		},
	}

	// Create a flow
	flow := &models.Flow{
		Name: "Test Flow",
		Steps: []models.Step{
			{
				ID:     "1",
				Name:   "Mock Step",
				Action: "mock/action",
				Inputs: map[string]interface{}{
					"input": "test",
				},
			},
		},
	}

	err := actions.RegisterAction(mockAction)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Create a runtime
	runtime := NewRuntime()
	runtime.SetFlow(flow)

	// Run the flow
	err = runtime.Run(nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check if the mock action was executed
	if !mockAction.executed {
		t.Error("Mock action was not executed")
	}
}

// MockAction is a mock implementation of the Action interface for testing
type MockAction struct {
	identifier  string
	executeFunc func(ctx models.Context, inputs map[string]interface{}) ([]models.Output, error)
	executed    bool
}

func (m *MockAction) GetIdentifier() string {
	return m.identifier
}

func (m *MockAction) GetInputSchema() []models.Input {
	return []models.Input{}
}

func (m *MockAction) GetOutputSchema() []models.Output {
	return []models.Output{}
}

func (m *MockAction) Validate(inputs map[string]interface{}) error {
	return nil
}

func (m *MockAction) Execute(ctx models.Context, inputs map[string]interface{}) ([]models.Output, error) {
	m.executed = true
	return m.executeFunc(ctx, inputs)
}
