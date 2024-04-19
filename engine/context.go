package engine

// Context represents the context of the flow execution and is passed to the actions (shared context between actions)
type Context struct {
	ExecutionID string

	Inputs  map[string]string
	Outputs map[string]string

	Logger Logger
}

func NewContext(executionID, flowName, flowVersion string, inputs map[string]string, logger Logger) *Context {
	return &Context{
		ExecutionID: executionID,
		Inputs:      inputs,
		Outputs:     make(map[string]string),
		Logger:      logger,
	}
}
