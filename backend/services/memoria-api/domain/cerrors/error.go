package cerrors

import "fmt"

type Validation struct {
	message string
}

func NewValidation(message string) error {
	return &Validation{message: message}
}

func (e Validation) Error() string {
	return e.message
}

type ResourceNotFound struct {
	message string
}

func NewResourceNotFound(message string) error {
	return &ResourceNotFound{message: message}
}

func (e ResourceNotFound) Error() string {
	return fmt.Sprintf("resource not found: %s", e.message)
}

type Internal struct {
	message string
}

func NewInternal(message string) error {
	return &Internal{message: message}
}

func (e Internal) Error() string {
	return "internal error"
}

type Unauthorized struct{}

func NewUnauthorized() error {
	return &Unauthorized{}
}

func (e Unauthorized) Error() string {
	return "unauthorized"
}
