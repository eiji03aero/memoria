package cerrors

import "fmt"

// -------------------- Consistency --------------------
type Consistency struct {
	message string
}

func NewConsistency(message string) error {
	return Consistency{message: message}
}

func (e Consistency) Error() string {
	return fmt.Sprintf("consistency error: %s", e.message)
}

// -------------------- Resource not found --------------------
type ResourceNotFound struct {
	name string
}

func NewResourceNotFound(name string) error {
	return ResourceNotFound{name: name}
}

func (e ResourceNotFound) Error() string {
	return fmt.Sprintf("resource not found: %s", e.name)
}

// -------------------- Internal --------------------
type Internal struct {
	message string
}

func NewInternal(message string) error {
	return Internal{message: message}
}

func (e Internal) Error() string {
	return fmt.Sprintf("internal error: %s", e.message)
}

// -------------------- Unauthorized --------------------
type Unauthorized struct{}

func NewUnauthorized() error {
	return Unauthorized{}
}

func (e Unauthorized) Error() string {
	return "unauthorized"
}

// -------------------- Not implemented --------------------
type NotImplemented struct {
	Name string
}

func NewNotImplemented(name string) error {
	return NotImplemented{Name: name}
}

func (e NotImplemented) Error() string {
	return "not implemented: " + e.Name
}
