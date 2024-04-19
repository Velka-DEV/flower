package engine

import (
	"flower/models"
)

type Runtime struct {
	flow    *models.Flow
	context *Context
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) SetFlow(flow *models.Flow) {
	r.flow = flow
}

func (r *Runtime) Run() error {
	
}
