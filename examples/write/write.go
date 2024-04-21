package main

import (
	models2 "flower/internal/models"
	"github.com/goccy/go-yaml"
	"os"
)

func main() {
	flow := models2.Flow{
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
		Steps: []models2.Step{
			{
				ID:     "1",
				Name:   "Test regex",
				Action: "core/regex",
				Inputs: map[string]interface{}{
					"regex": "(\\d+)",
					"text":  "The number is 42.",
				},
			},
			{
				ID:     "2",
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
