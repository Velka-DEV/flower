package flower

import "flower/internal/actions"

func RegisterAction(action Action) error {
	return actions.RegisterAction(action)
}

func RegisterActions(_actions []Action) error {
	return actions.RegisterActions(_actions)
}

func GetAction(identifier string) (Action, bool) {
	return actions.GetAction(identifier)
}

func GetActions() map[string]Action {
	return actions.GetActions()
}
