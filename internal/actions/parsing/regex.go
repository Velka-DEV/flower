package parsing

import (
	models2 "flower/internal/models"
	"regexp"
)

type RegexAction struct{}

func (a *RegexAction) GetIdentifier() string {
	return "parsing/regex"
}

func (a *RegexAction) GetInputSchema() []models2.Input {
	return []models2.Input{
		{
			Name:        "regex",
			Description: "The regular expression to match against the text.",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "text",
			Description: "The text to match against the regular expression.",
			Type:        "string",
			Required:    true,
		},
	}
}

func (a *RegexAction) GetOutputSchema() []models2.Output {
	return []models2.Output{
		{
			Name: "matches",
			Type: "array",
		},
		{
			Name: "groups",
			Type: "array",
		},
	}
}

func (a *RegexAction) Validate(inputs map[string]interface{}) error {
	if _, ok := inputs["regex"].(string); !ok {
		return &models2.InputValidationError{InputName: "regex", Message: "must be a string"}
	}
	if _, ok := inputs["text"].(string); !ok {
		return &models2.InputValidationError{InputName: "text", Message: "must be a string"}
	}
	return nil
}

func (a *RegexAction) Execute(ctx models2.Context, inputs map[string]interface{}) ([]models2.Output, error) {
	regex := inputs["regex"].(string)
	text := inputs["text"].(string)

	compiled, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	matches := compiled.FindAllString(text, -1)
	var groups [][]string
	if len(matches) > 0 {
		for _, match := range matches {
			groups = append(groups, compiled.FindStringSubmatch(match))
		}
	}

	return []models2.Output{
		{
			Name:  "matches",
			Value: matches,
		},
		{
			Name:  "groups",
			Value: groups,
		},
	}, nil
}
