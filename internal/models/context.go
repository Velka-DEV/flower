package models

type Context interface {
	SetOutput(stepID, outputName string, value interface{})
}
