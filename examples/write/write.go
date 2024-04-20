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
			"test_string": "Hello, world!",
			"test_int":    42,
			"test_bool":   true,
			"test_array":  []string{"a", "b", "c"},
			"test_object": map[string]interface{}{
				"key": "value",
			},
		},
		Steps: []models.Step{
			{
				Name:   "Test regex",
				Action: "core/regex",
				Inputs: map[string]interface{}{
					"regex": "(\\d+)",
					"text":  "The number is 42.",
				},
			},
			{
				Name:   "Test print",
				Action: "core/test/print",
				Inputs: map[string]interface{}{
					"message": "The matched number is {{index .matches 0}}",
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
