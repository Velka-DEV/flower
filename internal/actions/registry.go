package actions

import (
	core2 "flower/internal/actions/core"
	"flower/internal/actions/parsing"
	models2 "flower/internal/models"
)

var actions = []models2.Action{
	&core2.PrintAction{},
	&parsing.RegexAction{},
}

func RegisterAction(action models2.Action) error {
	if _, ok := GetAction(action.GetIdentifier()); ok {
		return &models2.ActionAlreadyRegisteredError{
			Action: action.GetIdentifier(),
		}
	}

	actions = append(actions, action)

	return nil
}

func RegisterActions(actions []models2.Action) error {
	for _, action := range actions {
		err := RegisterAction(action)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetAction(identifier string) (models2.Action, bool) {
	for _, action := range actions {
		if action.GetIdentifier() == identifier {
			return action, true
		}
	}

	return nil, false
}

func GetActions() map[string]models2.Action {
	actions := make(map[string]models2.Action)

	for _, action := range actions {
		actions[action.GetIdentifier()] = action
	}

	return actions
}
