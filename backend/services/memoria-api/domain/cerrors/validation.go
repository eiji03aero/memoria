package cerrors

import "fmt"

type ValidationKey string

var (
	ValidationKey_Required         = ValidationKey("required")
	ValidationKey_Invalid          = ValidationKey("invalid")
	ValidationKey_InvalidFormat    = ValidationKey("invalid-format")
	ValidationKey_AlreadyTaken     = ValidationKey("already-taken")
	ValidationKey_ResourceNotFound = ValidationKey("resource-not-found")
)

type Validation struct {
	Key  ValidationKey
	Name string
}

type NewValidationDTO struct {
	Key  ValidationKey
	Name string
}

func NewValidation(dto NewValidationDTO) error {
	return Validation{Key: dto.Key, Name: dto.Name}
}

func (e Validation) Error() string {
	return fmt.Sprintf("validation error: %s %s", e.Key, e.Name)
}
