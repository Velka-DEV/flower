package parsing

import (
	models2 "flower/internal/models"
	"strings"
)

type DelimiterParseAction struct{}

func (a *DelimiterParseAction) GetIdentifier() string {
	return "parsing/delimiter"
}

func (a *DelimiterParseAction) GetInputSchema() []models2.Input {
	return []models2.Input{
		{
			Name:        "text",
			Description: "The text to parse",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "start_delimiter",
			Description: "The starting delimiter",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "end_delimiter",
			Description: "The ending delimiter",
			Type:        "string",
			Required:    true,
		},
	}
}

func (a *DelimiterParseAction) GetOutputSchema() []models2.Output {
	return []models2.Output{
		{
			Name: "result",
			Type: "[]string",
		},
	}
}

func (a *DelimiterParseAction) Validate(inputs map[string]interface{}) error {
	if _, ok := inputs["text"].(string); !ok {
		return &models2.InputValidationError{InputName: "text", Message: "must be a string"}
	}
	if _, ok := inputs["start_delimiter"].(string); !ok {
		return &models2.InputValidationError{InputName: "start_delimiter", Message: "must be a string"}
	}
	if _, ok := inputs["end_delimiter"].(string); !ok {
		return &models2.InputValidationError{InputName: "end_delimiter", Message: "must be a string"}
	}
	return nil
}

func (a *DelimiterParseAction) Execute(ctx models2.Context, inputs map[string]interface{}) ([]models2.Output, error) {
	text := inputs["text"].(string)
	startDelimiter := inputs["start_delimiter"].(string)
	endDelimiter := inputs["end_delimiter"].(string)

	var result []string
	startIndex := 0

	for {
		// Find the start delimiter
		startPos := strings.Index(text[startIndex:], startDelimiter)
		if startPos == -1 {
			break
		}
		startPos += startIndex

		// Find the end delimiter
		endPos := strings.Index(text[startPos+len(startDelimiter):], endDelimiter)
		if endPos == -1 {
			break
		}
		endPos += startPos + len(startDelimiter)

		// Extract the text between delimiters
		extracted := text[startPos+len(startDelimiter) : endPos]
		result = append(result, extracted)

		// Move the start index for the next iteration
		startIndex = endPos + len(endDelimiter)
	}

	return []models2.Output{
		{
			Name:  "result",
			Value: result,
		},
	}, nil
}
