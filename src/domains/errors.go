package domains

import "fmt"

type DatabaseError struct {
	Message string
	Err     error
}

func (d DatabaseError) Error() string {
	if d.Err != nil {
		return fmt.Sprintf("%s. err: %s", d.Message, d.Err.Error())
	}
	return fmt.Sprintf("%s", d.Message)
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
