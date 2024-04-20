package engine

import (
	"flower/actions"
	"flower/models"
	"time"
)

type Runtime struct {
	flow    *models.Flow
	actions map[string]models.Action
	context *Context
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) SetFlow(flow *models.Flow) {
	r.flow = flow

	requiredActions, err := r.getRequiredActions()

	if err != nil {
		panic(err)
	}

	r.actions = requiredActions
}

func (r *Runtime) SetActions(actions map[string]models.Action) {
	r.actions = actions
}

func (r *Runtime) Run() error {
	r.context = NewContext(r.flow)
	startTime := time.Now()

	defer func() {
		r.context.Logger.Info("Execution finished in %v", time.Since(startTime))
	}()

	for _, step := range r.flow.Steps {
		stepStartTime := time.Now()

		action, ok := r.actions[step.Action]

		if !ok {
			r.context.Logger.Error("Action %s not found", step.Action)

			return &models.ActionNotFoundError{
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
			r.context.SetOutput(output.Name, output.Value)
		}

		r.context.Logger.Info("Step %s executed in %v", step.Name, time.Since(stepStartTime))
	}

	return nil
}

func (r *Runtime) getRequiredActions() (map[string]models.Action, error) {
	requiredActions := make(map[string]models.Action)

	for _, step := range r.flow.Steps {
		action, ok := actions.GetAction(step.Action)

		if !ok {
			return nil, &models.ActionNotFoundError{
				Action:  step.Action,
				Message: "Action not found",
			}
		}

		requiredActions[action.GetIdentifier()] = action
	}

	return requiredActions, nil
}
