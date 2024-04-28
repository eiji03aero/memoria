package model

type UserSpace struct {
	ID   string
	Name string
}

type NewUserSpaceDTO struct {
	ID   string
	Name string
}

func NewUserSpace(dto NewUserSpaceDTO) *UserSpace {
	return &UserSpace{
		ID:   dto.ID,
		Name: dto.Name,
	}
}
