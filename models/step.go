package models

type Step struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Action string `json:"action"`

	Inputs map[string]interface{} `json:"inputs"`
}
