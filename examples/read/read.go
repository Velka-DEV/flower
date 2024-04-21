package main

import (
	"flower/engine"
	"flower/models"
	"os"
)

func main() {
	data, err := os.ReadFile("flow.yaml")

	if err != nil {
		panic(err)
	}

	flow, err := models.FromYaml(data)

	if err != nil {
		panic(err)
	}

	runtime := engine.NewRuntime()

	runtime.SetFlow(flow)

	err = runtime.Run(nil)

	if err != nil {
		panic(err)
	}
}
