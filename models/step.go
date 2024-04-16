package models

type Step struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`

	Inputs  []Input
	Outputs []Output
}
