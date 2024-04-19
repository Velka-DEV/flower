package models

type Output struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Type        string      `json:"type"`
	Value       interface{} `json:"value"`
}
