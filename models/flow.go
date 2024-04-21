package models

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"regexp"
)

type Flow struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	FlowVersion string `json:"flow_version"`

	Inputs map[string]interface{} `json:"inputs"`

	Steps []Step `json:"steps"`
}

func (f *Flow) ToYaml() ([]byte, error) {
	return yaml.Marshal(f)
}

func FromYaml(data []byte) (*Flow, error) {
	flow := &Flow{}
	err := yaml.Unmarshal(data, flow)

	if err != nil {
		return nil, err
	}

	if flow.Inputs == nil {
		flow.Inputs = make(map[string]interface{})
	}

	if flow.Steps == nil {
		return nil, fmt.Errorf("flow must have steps")
	}

	if len(flow.Steps) == 0 {
		return nil, fmt.Errorf("flow must have at least one step")
	}

	if flow.FlowVersion == "" {
		return nil, fmt.Errorf("flow must have a version")
	}

	for i := range flow.Steps {
		if flow.Steps[i].ID == "" {
			return nil, fmt.Errorf("all steps must have an ID")
		}

		regex := "^[a-zA-Z0-9_]+$"

		if !regexp.MustCompile(regex).MatchString(flow.Steps[i].ID) {
			return nil, fmt.Errorf("step ID must match regex %s", regex)
		}
	}

	return flow, err
}
