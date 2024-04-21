# Flower

Flower is a powerful and flexible Go library for creating and executing workflows. It provides a simple and intuitive way to define workflows using a declarative YAML syntax and execute them programmatically.

## Features

- Define workflows using a simple and expressive YAML syntax
- Execute workflows programmatically using the Flower engine
- Built-in support for common actions (e.g., print, sleep, http)
- Extensible architecture for adding custom actions
- Comprehensive error handling and logging capabilities
- Lightweight and easy to integrate into existing Go projects

## Built-in Actions

⚠️ **Note (WIP):** This list is not exhaustive and may change over time. Please refer to the [actions](pkg/actions/) directory for the most up-to-date list of built-in actions.

Flower includes a number of built-in actions that can be used in workflows:

- `core/test/print`: Print a message to the console
- `core/regex`: Perform a regular expression match

## Installation

To install the Flower library, use the following command:

```
go get github.com/Velka-DEV/flower
```

## Usage

Here's a simple example of how to use Flower to define and execute a workflow:

```go
package main

import (
    "fmt"
    "github.com/Velka-DEV/flower/pkg/flower"
)

func main() {
    // Define the workflow using YAML syntax
    workflowYAML := `
name: Example Workflow
steps:
  - id: step1
    name: Print Hello
    action: core/print
    inputs:
      message: Hello, World!
`

    // Parse the workflow YAML
    workflow, err := flower.FromYaml([]byte(workflowYAML))
    if err != nil {
        panic(err)
    }

    // Create a new Flower runtime
    runtime := flower.NewRuntime()

    // Set the workflow in the runtime
    runtime.SetFlow(workflow)

    // Execute the workflow
    err = runtime.Run(nil)
    if err != nil {
        panic(err)
    }
}
```

For more detailed examples and usage instructions, please refer to the [examples](examples/) directory.

## Contributing

Contributions are welcome! If you'd like to contribute to Flower, please follow these steps:

1. Fork the repository
2. Create a new branch for your feature or bug fix
3. Make your changes and commit them with descriptive commit messages
4. Push your changes to your forked repository
5. Submit a pull request to the main repository

Please ensure that your code follows the existing coding style and includes appropriate tests.

## License

Flower is released under the [GPLv3 License](LICENSE).

## Contact

If you have any questions, suggestions, or feedback, please feel free to contact us at:

- Email: contact@velka.dev
- GitHub Issues: [https://github.com/Velka-DEV/flower/issues](https://github.com/Velka-DEV/flower/issues)

We appreciate your interest in Flower and look forward to hearing from you!