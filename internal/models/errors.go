package models

type ActionNotFoundError struct {
	Message string
	Action  string
}

func (e *ActionNotFoundError) Error() string {
	return "Action not found: " + e.Action + " - " + e.Message
}

type InputValidationError struct {
	Message   string
	InputName string
}

func (e *InputValidationError) Error() string {
	return "Input validation error: " + e.InputName + " - " + e.Message
}

type ActionAlreadyRegisteredError struct {
	Action string
}

func (e *ActionAlreadyRegisteredError) Error() string {
	return "Action already registered: " + e.Action
}
