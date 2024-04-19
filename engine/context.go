package engine

import (
	"flower/models"
	"github.com/google/uuid"
	"time"
)

// Context represents the context of the flow execution and is passed to the actions (shared context between actions)
type Context struct {
	ExecutionID string
	Inputs      map[string]string
	Logger      Logger

	startTime time.Time
	endTime   time.Time

	flow        *models.Flow
	stepOutputs map[string]interface{}
}

func NewContext(flow *models.Flow) *Context {
	id, _ := uuid.NewUUID()

	return &Context{
		ExecutionID: id.String(),
		Inputs:      make(map[string]string),
		flow:        flow,
		Logger:      NewDefaultLogger(id.String()),
	}
}

func (c *Context) SetOutput(name string, value interface{}) {
	c.stepOutputs[name] = value
}

func (c *Context) GetOutput(name string) (interface{}, bool) {
	value, ok := c.stepOutputs[name]
	return value, ok
}
