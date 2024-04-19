package main

import (
	"flower/engine"
	"flower/models"
)

func main() {
	flow := models.Flow{
		Name:        "Test Flow",
		Description: "Testing simple print message action in a flow",
		Version:     "0.0.1",
		Author:      "Flower",
		FlowVersion: "0.0.1",
		Steps: []models.Step{
			{
				Name:   "Test print",
				Action: "core/test/print",

				Inputs: map[string]interface{}{
					"message": "Hello, world!",
				},
			},
		},
	}

	runtime := engine.NewRuntime()

	runtime.SetFlow(&flow)

	err := runtime.Run()

	if err != nil {
		panic(err)
	}
}
