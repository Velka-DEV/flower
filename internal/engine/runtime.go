package engine

import (
	models2 "flower/internal/models"
	"time"
)

type Runtime struct {
	flow    *models2.Flow
	actions map[string]models2.Action
	context *Context
}

func NewRuntime() *Runtime {
	return &Runtime{
		actions: make(map[string]models2.Action),
	}
}

func (r *Runtime) SetFlow(flow *models2.Flow) {
	r.flow = flow
}

func (r *Runtime) SetActions(actions map[string]models2.Action) {
	r.actions = actions
}

func (r *Runtime) GetActions() map[string]models2.Action {
	return r.actions
}

func (r *Runtime) GetFlow() *models2.Flow {
	return r.flow
}

func (r *Runtime) GetContext() *Context {
	return r.context
}

func (r *Runtime) Run(inputs map[string]interface{}) error {
	if err := r.validateActions(); err != nil {
		return err
	}

	r.context = NewContext(r.flow, inputs)
	startTime := time.Now()

	defer func() {
		r.context.Logger.Info("Execution finished in %v", time.Since(startTime))
	}()

	for _, step := range r.flow.Steps {
		stepStartTime := time.Now()

		action, ok := r.actions[step.Action]
		if !ok {
			r.context.Logger.Error("Action %s not found", step.Action)
			return &models2.ActionNotFoundError{
				Action:  step.Action,
				Message: "Action not found",
			}
		}

		resolvedInputs, err := resolveInputs(r.context, step.Inputs)
		if err != nil {
			r.context.Logger.Error("Error resolving inputs for step %s: %v", step.Name, err)
			return err
		}

		err = action.Validate(resolvedInputs)
		if err != nil {
			r.context.Logger.Error("Validation error for step %s: %v", step.Name, err)
			return err
		}

		outputs, err := action.Execute(r.context, resolvedInputs)
		if err != nil {
			r.context.Logger.Error("Error executing step %s: %v", step.Name, err)
			return err
		}

		for _, output := range outputs {
			r.context.SetOutput(step.ID, output.Name, output.Value)
		}

		r.context.Logger.Info("Step %s executed in %v", step.Name, time.Since(stepStartTime))
	}

	return nil
}

func (r *Runtime) validateActions() error {
	for _, step := range r.flow.Steps {
		if _, ok := r.actions[step.Action]; !ok {
			return &models2.ActionNotFoundError{
				Action:  step.Action,
				Message: "Action not found",
			}
		}
	}
	return nil
}
