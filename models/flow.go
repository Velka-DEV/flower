package models

type Flow struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`

	Inputs []Input
	Tasks  []Task
}
