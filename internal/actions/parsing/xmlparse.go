package parsing

import (
	models2 "flower/internal/models"
	"fmt"
	"github.com/clbanning/mxj/v2"
)

type XMLParseAction struct{}

func (a *XMLParseAction) GetIdentifier() string {
	return "parsing/xml"
}

func (a *XMLParseAction) GetInputSchema() []models2.Input {
	return []models2.Input{
		{
			Name:        "xml",
			Description: "The XML string to parse",
			Type:        "string",
			Required:    true,
		},
	}
}

func (a *XMLParseAction) GetOutputSchema() []models2.Output {
	return []models2.Output{
		{
			Name: "result",
			Type: "map[string]interface{}",
		},
	}
}

func (a *XMLParseAction) Validate(inputs map[string]interface{}) error {
	if _, ok := inputs["xml"].(string); !ok {
		return &models2.InputValidationError{InputName: "xml", Message: "must be a string"}
	}
	return nil
}

func (a *XMLParseAction) Execute(ctx models2.Context, inputs map[string]interface{}) ([]models2.Output, error) {
	xmlString := inputs["xml"].(string)

	mv, err := mxj.NewMapXml([]byte(xmlString))
	if err != nil {
		return nil, fmt.Errorf("error parsing XML: %v", err)
	}

	// Convert mxj.Map to map[string]interface{}
	result := map[string]interface{}(mv)

	return []models2.Output{
		{
			Name:  "result",
			Value: result,
		},
	}, nil
}
