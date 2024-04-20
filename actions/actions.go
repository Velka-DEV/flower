package actions

import (
	"flower/actions/core"
	"flower/models"
)

var actions = []models.Action{
	&core.PrintAction{},
	&core.RegexAction{},
}

func RegisterAction(action models.Action) error {
	if _, ok := GetAction(action.GetIdentifier()); ok {
		return &models.ActionAlreadyRegisteredError{
			Action: action.GetIdentifier(),
		}
	}

	actions = append(actions, action)

	return nil
}

func RegisterActions(actions []models.Action) error {
	for _, action := range actions {
		err := RegisterAction(action)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetAction(identifier string) (models.Action, bool) {
	for _, action := range actions {
		if action.GetIdentifier() == identifier {
			return action, true
		}
	}

	return nil, false
}

func GetActions() map[string]models.Action {
	actions := make(map[string]models.Action)

	for _, action := range actions {
		actions[action.GetIdentifier()] = action
	}

	return actions
}
