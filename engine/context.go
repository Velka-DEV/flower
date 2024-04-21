package engine

import (
	"flower/models"
	"github.com/google/uuid"
	"time"
)

// Context represents the context of the flow execution and is passed to the actions (shared context between actions)
type Context struct {
	ExecutionID uuid.UUID
	Inputs      map[string]interface{}
	Logger      Logger

	startTime time.Time
	endTime   time.Time

	flow        *models.Flow
	stepOutputs map[string]map[string]interface{}
}

func NewContext(flow *models.Flow, inputs map[string]interface{}) *Context {
	id, _ := uuid.NewUUID()

	if inputs == nil {
		inputs = make(map[string]interface{})
	}

	for inputName, input := range flow.Inputs {
		if _, ok := inputs[inputName]; !ok {
			inputs[inputName] = input
		}
	}

	return &Context{
		ExecutionID: id,
		Inputs:      inputs,
		flow:        flow,
		Logger:      NewDefaultLogger(id),
		stepOutputs: make(map[string]map[string]interface{}),
	}
}

func (c *Context) SetOutput(stepID, outputName string, value interface{}) {
	if _, ok := c.stepOutputs[stepID]; !ok {
		c.stepOutputs[stepID] = make(map[string]interface{})
	}
	c.stepOutputs[stepID][outputName] = value
}

func (c *Context) GetOutput(stepID, outputName string) (interface{}, bool) {
	stepOutputs, ok := c.stepOutputs[stepID]
	if !ok {
		return nil, false
	}
	value, ok := stepOutputs[outputName]
	return value, ok
}
