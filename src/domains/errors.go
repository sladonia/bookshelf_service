package domains

import "fmt"

type DatabaseError struct {
	Message string
	Err     error
}

func (d DatabaseError) Error() string {
	return fmt.Sprintf("%s. err: %s", d.Message, d.Err.Error())
}

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

func NewValidationError(message string) ValidationError {
	return ValidationError{Message: message}
}
