package models

type Flow struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	FlowVersion string `json:"flow_version"`

	Inputs map[string]interface{} `json:"inputs"`

	Steps []Step `json:"steps"`
}
