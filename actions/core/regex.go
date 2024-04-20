package core

import (
	"flower/models"
	"fmt"
	"regexp"
)

const RegexActionIdentifier = "core/regex"

type RegexAction struct {
}

func (a *RegexAction) GetIdentifier() string {
	return RegexActionIdentifier
}

func (a *RegexAction) GetInputSchema() []models.Input {
	return []models.Input{
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

func (a *RegexAction) GetOutputSchema() []models.Output {
	return []models.Output{
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
	value, ok := inputs["regex"]

	if !ok {
		return fmt.Errorf("regex is required")
	}

	if _, ok := value.(string); !ok {
		return fmt.Errorf("regex must be a string")
	}

	return nil
}

func (a *RegexAction) Execute(ctx models.Context, inputs map[string]interface{}) ([]models.Output, error) {
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

	outputs := []models.Output{
		{
			Name:  "matches",
			Value: matches,
		},
		{
			Name:  "groups",
			Value: groups,
		},
	}

	return outputs, nil
}
