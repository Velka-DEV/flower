package requests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPRequestAction_Execute(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		// Check the request headers
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header 'application/json', got '%s'", r.Header.Get("Accept"))
		}

		// Send response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello, World!"}`))
	}))
	defer server.Close()

	// Create an instance of HTTPRequestAction
	action := &HTTPRequestAction{}

	// Create test inputs
	inputs := map[string]interface{}{
		"url":    server.URL,
		"method": "GET",
		"headers": map[string]string{
			"Accept": "application/json",
		},
	}

	// Execute the action
	outputs, err := action.Execute(nil, inputs)

	// Check for errors
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check outputs
	if len(outputs) != 3 {
		t.Fatalf("Expected 3 outputs, got %d", len(outputs))
	}

	// Check status code
	statusCode, ok := outputs[0].Value.(int)
	if !ok || statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %v", http.StatusOK, outputs[0].Value)
	}

	// Check response body
	body, ok := outputs[2].Value.(string)
	if !ok || body != `{"message": "Hello, World!"}` {
		t.Errorf("Unexpected response body: %v", outputs[2].Value)
	}
}

func TestHTTPRequestAction_Validate(t *testing.T) {
	action := &HTTPRequestAction{}

	tests := []struct {
		name    string
		inputs  map[string]interface{}
		wantErr bool
	}{
		{
			name: "Valid inputs",
			inputs: map[string]interface{}{
				"url":    "https://example.com",
				"method": "GET",
			},
			wantErr: false,
		},
		{
			name: "Missing URL",
			inputs: map[string]interface{}{
				"method": "GET",
			},
			wantErr: true,
		},
		{
			name: "Invalid method type",
			inputs: map[string]interface{}{
				"url":    "https://example.com",
				"method": 123,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := action.Validate(tt.inputs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
