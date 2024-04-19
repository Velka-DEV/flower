package main

import (
	"flower/models"
	yaml "github.com/goccy/go-yaml"
	"os"
)

func main() {
	flow := models.Flow{
		Name:        "Test Flow",
		Description: "Testing object to yaml export",
		Version:     "0.0.1",
		Author:      "Flower",
		FlowVersion: "0.0.1",
		Inputs: map[string]interface{}{
			"test":  "Testing input",
			"test2": "Testing input 2",
		},
		Steps: []models.Step{
			models.Step{
				Name:   "Test Step",
				Type:   "test",
				Action: "request/http",

				Inputs: []models.Input{
					models.Input{
						Name:        "Test Input",
						Description: "Testing input",
						Type:        "string",
						Required:    true,
						Default:     "default",
						Options:     []string{"option1", "option2"},
						Value:       "value",
					},
				},
			},
		},
	}

	yaml, err := yaml.Marshal(flow)

	if err != nil {
		panic(err)
	}

	// Write the yaml to a file
	err = os.WriteFile("flow.yaml", yaml, 0644)
}
