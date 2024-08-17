package request

import (
	"bytes"
	models2 "flower/internal/models"
	"fmt"
	"io"
	"net/http"
	"time"
)

const HTTPRequestActionIdentifier = "request/http"

type HTTPRequestAction struct{}

func (a *HTTPRequestAction) GetIdentifier() string {
	return HTTPRequestActionIdentifier
}

func (a *HTTPRequestAction) GetInputSchema() []models2.Input {
	return []models2.Input{
		{
			Name:        "url",
			Description: "The URL to send the request to",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "method",
			Description: "The HTTP method to use (GET, POST, PUT, DELETE, etc.)",
			Type:        "string",
			Required:    true,
		},
		{
			Name:        "headers",
			Description: "A map of headers to send with the request",
			Type:        "map[string]string",
			Required:    false,
		},
		{
			Name:        "body",
			Description: "The body of the request (for POST, PUT, etc.)",
			Type:        "string",
			Required:    false,
		},
		{
			Name:        "timeout",
			Description: "Request timeout in seconds",
			Type:        "int",
			Required:    false,
			Default:     "30",
		},
	}
}

func (a *HTTPRequestAction) GetOutputSchema() []models2.Output {
	return []models2.Output{
		{
			Name: "status_code",
			Type: "int",
		},
		{
			Name: "headers",
			Type: "map[string][]string",
		},
		{
			Name: "body",
			Type: "string",
		},
	}
}

func (a *HTTPRequestAction) Validate(inputs map[string]interface{}) error {
	if _, ok := inputs["url"].(string); !ok {
		return &models2.InputValidationError{InputName: "url", Message: "must be a string"}
	}
	if _, ok := inputs["method"].(string); !ok {
		return &models2.InputValidationError{InputName: "method", Message: "must be a string"}
	}
	if headers, ok := inputs["headers"]; ok {
		if _, ok := headers.(map[string]string); !ok {
			return &models2.InputValidationError{InputName: "headers", Message: "must be a map[string]string"}
		}
	}
	if timeout, ok := inputs["timeout"]; ok {
		if _, ok := timeout.(int); !ok {
			return &models2.InputValidationError{InputName: "timeout", Message: "must be an integer"}
		}
	}
	return nil
}

func (a *HTTPRequestAction) Execute(ctx models2.Context, inputs map[string]interface{}) ([]models2.Output, error) {
	url := inputs["url"].(string)
	method := inputs["method"].(string)

	var headers map[string]string
	if h, ok := inputs["headers"]; ok {
		headers = h.(map[string]string)
	}

	var body string
	if b, ok := inputs["body"]; ok {
		body = b.(string)
	}

	timeout := 30
	if t, ok := inputs["timeout"]; ok {
		timeout = t.(int)
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	var reqBody io.Reader
	if body != "" {
		reqBody = bytes.NewBufferString(body)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	outputs := []models2.Output{
		{
			Name:  "status_code",
			Value: resp.StatusCode,
		},
		{
			Name:  "headers",
			Value: resp.Header,
		},
		{
			Name:  "body",
			Value: string(respBody),
		},
	}

	return outputs, nil
}
