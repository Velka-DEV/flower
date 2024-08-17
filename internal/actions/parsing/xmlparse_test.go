package parsing

import (
	"reflect"
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
		t.Errorf("Unexpected output structure: %v", outputs)
	}

	result, ok := outputs[0].Value.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", outputs[0].Value)
	}

	expected := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "John",
			"age":  "30",
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestXMLParseAction_ExecuteComplex(t *testing.T) {
	action := &XMLParseAction{}
	inputs := map[string]interface{}{
		"xml": `
		<root>
			<person>
				<name>John</name>
				<age>30</age>
				<address>
					<street>123 Main St</street>
					<city>New York</city>
				</address>
			</person>
			<person>
				<name>Jane</name>
				<age>28</age>
				<address>
					<street>456 Elm St</street>
					<city>Los Angeles</city>
				</address>
			</person>
		</root>`,
	}

	outputs, err := action.Execute(nil, inputs)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result, ok := outputs[0].Value.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", outputs[0].Value)
	}

	expected := map[string]interface{}{
		"root": map[string]interface{}{
			"person": []interface{}{
				map[string]interface{}{
					"name": "John",
					"age":  "30",
					"address": map[string]interface{}{
						"street": "123 Main St",
						"city":   "New York",
					},
				},
				map[string]interface{}{
					"name": "Jane",
					"age":  "28",
					"address": map[string]interface{}{
						"street": "456 Elm St",
						"city":   "Los Angeles",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
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
