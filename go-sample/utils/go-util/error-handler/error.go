package error_handler

import (
	"fmt"
	"reflect"
)

func NewDomainError(message string, code int, details interface{}) DomainError {
	return DomainError{
		Message: message,
		Code:    code,
		Details: details,
	}
}

func NewApplocationError(message string, details interface{}) ApplicationError {
	return ApplicationError{
		Message: message,
		Details: details,
	}
}

type DomainError struct {
	Message string
	Details interface{}
	Code    int
}

func (error DomainError) Error() string {
	return fmt.Sprintf(`%s | %v | %d`, error.Message, error.Details, error.Code)
}

type ApplicationError struct {
	Message string
	Details interface{}
}

func (error ApplicationError) Error() string {
	return fmt.Sprintf(`%s | %v`, error.Message, error.Details)
}

func IsDomain(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(DomainError{})
}
