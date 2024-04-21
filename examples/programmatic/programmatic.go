package main

import (
	"flower/engine"
	"flower/models"
)

func main() {
	flow := models.Flow{
		Name:        "Test Flow",
		Description: "Testing programmatic print message action in a flow",
		Version:     "0.0.1",
		Author:      "Flower",
		FlowVersion: "0.0.1",
		Steps: []models.Step{
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

	runtime := engine.NewRuntime()

	runtime.SetFlow(&flow)

	err := runtime.Run(nil)

	if err != nil {
		panic(err)
	}
}
