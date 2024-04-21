package models

type Step struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Action string `json:"action"`

	Inputs map[string]interface{} `json:"inputs"`
}
