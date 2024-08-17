package parsing

import (
	models2 "flower/internal/models"
	"fmt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"strings"
)

type JSONPathAction struct{}

func (a *JSONPathAction) GetIdentifier() string {
	return "parsing/jsonpath"
}

func (a *JSONPathAction) GetInputSchema() []models2.Input {
	return []models2.Input{
		{
			Name:        "json",
			Description: "The JSON string to parse",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "path",
			Description: "The JSONPath expression",
			Type:        "string",
			Required:    true,
		},
	}
}

func (a *JSONPathAction) GetOutputSchema() []models2.Output {
	return []models2.Output{
		{
			Name: "result",
			Type: "interface{}",
		},
	}
}

func (a *JSONPathAction) Validate(inputs map[string]interface{}) error {
	if _, ok := inputs["json"].(string); !ok {
		return &models2.InputValidationError{InputName: "json", Message: "must be a string"}
	}
	if _, ok := inputs["path"].(string); !ok {
		return &models2.InputValidationError{InputName: "path", Message: "must be a string"}
	}
	return nil
}

func (a *JSONPathAction) Execute(ctx models2.Context, inputs map[string]interface{}) ([]models2.Output, error) {
	jsonString := inputs["json"].(string)
	path := inputs["path"].(string)

	jsonData, err := oj.ParseString(jsonString)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	expr, err := jp.ParseString(path)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSONPath expression: %v", err)
	}

	result := expr.Get(jsonData)
	if len(result) == 0 {
		return []models2.Output{{Name: "result", Value: nil}}, nil
	}

	var outputValue interface{}
	if isMultipleItemExpression(path) {
		outputValue = result
	} else {
		outputValue = result[0]
	}

	return []models2.Output{
		{
			Name:  "result",
			Value: outputValue,
		},
	}, nil
}

func isMultipleItemExpression(path string) bool {
	// Check for wildcard or array slice notation
	return strings.Contains(path, "*") || strings.Contains(path, "[") && strings.Contains(path, ":")
}
