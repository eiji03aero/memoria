package svc

type Medium interface {
	Delete(dto MediumDeleteDTO) (err error)
}

type MediumDeleteDTO struct {
	MediumID string
}
