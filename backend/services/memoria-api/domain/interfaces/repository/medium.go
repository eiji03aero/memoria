package repository

import "memoria-api/domain/model"

type Medium interface {
	Find(findOpt *FindOption) (media []*model.Medium, err error)
	FindOne(findOpt *FindOption) (medium *model.Medium, err error)
	Create(dto MediumCreateDTO) (medium *model.Medium, err error)
	Delete(dto MediumDeleteDTO) (err error)
}

type MediumCreateDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	Name        string
	Extension   string
}

type MediumDeleteDTO struct {
	ID string
}
