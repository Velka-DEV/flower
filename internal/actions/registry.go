package actions

import (
	core2 "flower/internal/actions/core"
	"flower/internal/actions/parsing"
	"flower/internal/actions/request"
	models2 "flower/internal/models"
)

var registry = []models2.Action{
	// Core
	&core2.PrintAction{},

	// Requests
	&request.HTTPRequestAction{},

	// Parsing
	&parsing.RegexAction{},
	&parsing.JSONPathAction{},
	&parsing.XPathAction{},
	&parsing.XMLParseAction{},
}

func RegisterAction(action models2.Action) error {
	if _, ok := GetAction(action.GetIdentifier()); ok {
		return &models2.ActionAlreadyRegisteredError{
			Action: action.GetIdentifier(),
		}
	}

	registry = append(registry, action)

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
	for _, action := range registry {
		if action.GetIdentifier() == identifier {
			return action, true
		}
	}

	return nil, false
}

func GetActions() map[string]models2.Action {
	actions := make(map[string]models2.Action)

	for _, action := range registry {
		actions[action.GetIdentifier()] = action
	}

	return actions
}
