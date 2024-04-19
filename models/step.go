package models

type Step struct {
	Name   string `json:"name"`
	Action string `json:"action"`

	Inputs map[string]interface{} `json:"inputs"`
}
