package core

import (
	"flower/models"
	"fmt"
)

const PrintActionIdentifier = "core/test/print"

var printInputSchema = []models.Input{
	{
		Name:     "message",
		Type:     "string",
		Required: true,
	},
}

type PrintAction struct {
}

func (a *PrintAction) GetIdentifier() string {
	return PrintActionIdentifier
}

func (a *PrintAction) GetInputSchema() []models.Input {
	return printInputSchema
}

func (a *PrintAction) GetOutputSchema() []models.Output {
	return []models.Output{}
}

func (a *PrintAction) Validate(inputs map[string]interface{}) error {
	value, ok := inputs["message"]

	if !ok {
		return fmt.Errorf("message is required")
	}

	if _, ok := value.(string); !ok {
		return fmt.Errorf("message must be a string")
	}

	return nil
}

func (a *PrintAction) Execute(ctx models.Context, inputs map[string]interface{}) ([]models.Output, error) {
	message := inputs["message"].(string)
	fmt.Println(message)

	return []models.Output{}, nil
}
