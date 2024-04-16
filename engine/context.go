package engine

// Context represents the context of the flow execution and is passed to the actions (shared context between actions)
type Context struct {
	ExecutionID string
	FlowName    string
	FlowVersion string

	Inputs  map[string]string
	Outputs map[string]string

	Logger Logger
}
