package cerrors

import "fmt"

// -------------------- Validation --------------------
type Validation struct {
	message string
}

func NewValidation(message string) error {
	return &Validation{message: message}
}

func (e Validation) Error() string {
	return e.message
}

// -------------------- Consistency --------------------
type Consistency struct {
	message string
}

func NewConsistency(message string) error {
	return &Consistency{message: message}
}

func (e Consistency) Error() string {
	return e.message
}

// -------------------- Resource not found --------------------
type ResourceNotFound struct {
	message string
}

func NewResourceNotFound(message string) error {
	return &ResourceNotFound{message: message}
}

func (e ResourceNotFound) Error() string {
	return fmt.Sprintf("resource not found: %s", e.message)
}

// -------------------- Internal --------------------
type Internal struct {
	message string
}

func NewInternal(message string) error {
	return &Internal{message: message}
}

func (e Internal) Error() string {
	return "internal error"
}

// -------------------- Unauthorized --------------------
type Unauthorized struct{}

func NewUnauthorized() error {
	return &Unauthorized{}
}

func (e Unauthorized) Error() string {
	return "unauthorized"
}
