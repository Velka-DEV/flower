package parsing

import (
	"reflect"
	"testing"
)

func TestXPathAction_Execute(t *testing.T) {
	action := &XPathAction{}
	inputs := map[string]interface{}{
		"xml":   `<root><person><name>John</name><age>30</age></person><person><name>Jane</name><age>25</age></person></root>`,
		"xpath": "//person/name/text()",
	}

	outputs, err := action.Execute(nil, inputs)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(outputs) != 1 || outputs[0].Name != "result" {
		t.Errorf("Unexpected output: %v", outputs)
	}

	result, ok := outputs[0].Value.([]string)
	if !ok {
		t.Fatalf("Expected []string, got %T", outputs[0].Value)
	}

	expected := []string{"John", "Jane"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestXPathAction_Validate(t *testing.T) {
	action := &XPathAction{}

	validInputs := map[string]interface{}{
		"xml":   `<root><person><name>John</name></person></root>`,
		"xpath": "//person/name/text()",
	}

	if err := action.Validate(validInputs); err != nil {
		t.Errorf("Unexpected error for valid inputs: %v", err)
	}

	invalidInputs := map[string]interface{}{
		"xml":   123,
		"xpath": "//person/name/text()",
	}

	if err := action.Validate(invalidInputs); err == nil {
		t.Errorf("Expected error for invalid inputs, got nil")
	}
}
