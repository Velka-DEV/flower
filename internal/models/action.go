package models

type Action interface {
	GetIdentifier() string
	GetInputSchema() []Input
	GetOutputSchema() []Output
	Validate(inputs map[string]interface{}) error
	Execute(ctx Context, inputs map[string]interface{}) ([]Output, error)
}
