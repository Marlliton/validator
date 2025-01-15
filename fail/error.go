package fail

import "fmt"

type Error struct {
	Field   string
	Message string
}

func (v *Error) Error() string {
	return fmt.Sprintf("%s", v.Message)
}

func New(field, message string) *Error {
	return &Error{
		Field:   field,
		Message: message,
	}
}
