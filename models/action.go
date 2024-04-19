package models

import (
	"flower/engine"
)

type Action interface {
	GetInputSchema() []Input
	GetOutputSchema() []Output
	Validate(inputs map[string]interface{}) error
	Execute(ctx *engine.Context, inputs map[string]interface{}) ([]Output, error)
}
