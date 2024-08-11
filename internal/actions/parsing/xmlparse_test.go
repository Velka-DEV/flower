package parsing

import (
	"testing"
)

func TestXMLParseAction_Execute(t *testing.T) {
	action := &XMLParseAction{}
	inputs := map[string]interface{}{
		"xml": `<person><name>John</name><age>30</age></person>`,
	}

	outputs, err := action.Execute(nil, inputs)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(outputs) != 1 || outputs[0].Name != "result" {
		t.Errorf("Unexpected output: %v", outputs)
	}

	result, ok := outputs[0].Value.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", outputs[0].Value)
	}

	person, ok := result["person"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected person to be map[string]interface{}, got %T", result["person"])
	}

	if person["name"] != "John" || person["age"] != "30" {
		t.Errorf("Unexpected parsed XML: %v", person)
	}
}

func TestXMLParseAction_Validate(t *testing.T) {
	action := &XMLParseAction{}

	validInputs := map[string]interface{}{
		"xml": `<person><name>John</name></person>`,
	}

	if err := action.Validate(validInputs); err != nil {
		t.Errorf("Unexpected error for valid inputs: %v", err)
	}

	invalidInputs := map[string]interface{}{
		"xml": 123,
	}

	if err := action.Validate(invalidInputs); err == nil {
		t.Errorf("Expected error for invalid inputs, got nil")
	}
}
