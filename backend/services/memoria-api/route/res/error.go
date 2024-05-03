package res

import "memoria-api/domain/cerrors"

type ValidationRes struct {
	Validation struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"validation"`
}

func NewValidationRes(err cerrors.Validation) *ValidationRes {
	res := &ValidationRes{}
	res.Validation.Key = string(err.Key)
	res.Validation.Name = err.Name

	return res
}
