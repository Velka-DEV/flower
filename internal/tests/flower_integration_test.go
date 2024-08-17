package tests

import (
	"flower/internal/engine"
	"flower/internal/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFlowIntegration(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ip": "192.168.1.1", "country": "United States", "city": "New York"}`))
	}))
	defer server.Close()

	// Create a temporary YAML file for the flow configuration
	yamlContent := []byte(`
name: IP Information Flow
description: Fetch IP information and extract details
version: 1.0.0
author: Test Author
flow_version: 0.0.1
steps:
  - id: fetch_ip_info
    name: Fetch IP Information
    action: request/http
    inputs:
      url: ` + server.URL + `
      method: GET
  - id: extract_ip
    name: Extract IP Address
    action: parsing/jsonpath
    inputs:
      json: '{{index .steps "fetch_ip_info" "body"}}'
      path: $.ip
  - id: extract_country
    name: Extract Country
    action: parsing/jsonpath
    inputs:
      json: '{{index .steps "fetch_ip_info" "body"}}'
      path: $.country
  - id: validate_ip
    name: Validate IP Address
    action: parsing/regex
    inputs:
      regex: '^(\d{1,3}\.){3}\d{1,3}$'
      text: '{{index .steps "extract_ip" "result"}}'
`)

	tmpfile, err := ioutil.TempFile("", "flow-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(yamlContent); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Read the YAML file and create a Flow
	yamlData, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	flow, err := models.FromYaml(yamlData)
	if err != nil {
		t.Fatalf("Error creating flow from YAML: %v", err)
	}

	// Create a new runtime and set the flow
	runtime := engine.NewRuntime()
	runtime.SetFlow(flow)

	// Run the flow
	err = runtime.Run(nil)
	if err != nil {
		t.Fatalf("Error running flow: %v", err)
	}

	// Validate the results
	ctx := runtime.GetContext()

	ipResult, ok := ctx.GetOutput("extract_ip", "result")
	if !ok {
		t.Fatal("IP result not found")
	}
	if ip, ok := ipResult.(string); !ok || ip != "192.168.1.1" {
		t.Errorf("Expected IP 192.168.1.1, got %v", ip)
	}

	countryResult, ok := ctx.GetOutput("extract_country", "result")
	if !ok {
		t.Fatal("Country result not found")
	}
	if country, ok := countryResult.(string); !ok || country != "United States" {
		t.Errorf("Expected country United States, got %v", country)
	}

	ipValidationResult, ok := ctx.GetOutput("validate_ip", "matches")
	if !ok {
		t.Fatal("IP validation result not found")
	}
	if matches, ok := ipValidationResult.([]string); !ok || len(matches) != 1 || matches[0] != "192.168.1.1" {
		t.Errorf("Expected IP validation match [192.168.1.1], got %v", matches)
	}
}
