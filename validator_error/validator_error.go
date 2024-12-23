package validator_error

import "fmt"

type ValidatorError struct {
	Field   string
	Message string
}

func (v *ValidatorError) Error() string {
	return fmt.Sprintf("%s", v.Message)
}
