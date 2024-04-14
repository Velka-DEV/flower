package models

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Inputs  []Input
	Outputs []Output
}
