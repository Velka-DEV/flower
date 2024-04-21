package models

type Input struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type"`
	Required    bool     `json:"required"`
	Default     string   `json:"default,omitempty"`
	Options     []string `json:"options,omitempty"`

	// The input's value
	Value         interface{} `json:"value,omitempty"`
	ValueTemplate string      `json:"value_template,omitempty"`
}
