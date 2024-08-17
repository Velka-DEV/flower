package parsing

import (
	models2 "flower/internal/models"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"strings"
)

type XPathAction struct{}

func (a *XPathAction) GetIdentifier() string {
	return "parsing/xpath"
}

func (a *XPathAction) GetInputSchema() []models2.Input {
	return []models2.Input{
		{
			Name:        "xml",
			Description: "The XML string to parse",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "xpath",
			Description: "The XPath expression",
			Type:        "string",
			Required:    true,
		},
	}
}

func (a *XPathAction) GetOutputSchema() []models2.Output {
	return []models2.Output{
		{
			Name: "result",
			Type: "[]string",
		},
	}
}

func (a *XPathAction) Validate(inputs map[string]interface{}) error {
	if _, ok := inputs["xml"].(string); !ok {
		return &models2.InputValidationError{InputName: "xml", Message: "must be a string"}
	}
	if _, ok := inputs["xpath"].(string); !ok {
		return &models2.InputValidationError{InputName: "xpath", Message: "must be a string"}
	}
	return nil
}

func (a *XPathAction) Execute(ctx models2.Context, inputs map[string]interface{}) ([]models2.Output, error) {
	xmlString := inputs["xml"].(string)
	xpathExpr := inputs["xpath"].(string)

	doc, err := xmlquery.Parse(strings.NewReader(xmlString))
	if err != nil {
		return nil, fmt.Errorf("error parsing XML: %v", err)
	}

	expr, err := xpath.Compile(xpathExpr)
	if err != nil {
		return nil, fmt.Errorf("error compiling XPath expression: %v", err)
	}

	nodes, ok := expr.Evaluate(xmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
	if !ok {
		return nil, fmt.Errorf("unexpected result type from XPath evaluation")
	}

	var results []string
	for nodes.MoveNext() {
		results = append(results, nodes.Current().Value())
	}

	return []models2.Output{
		{
			Name:  "result",
			Value: results,
		},
	}, nil
}
