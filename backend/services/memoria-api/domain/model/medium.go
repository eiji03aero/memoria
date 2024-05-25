package model

import (
	"strings"
	"time"

	"memoria-api/config"
	"memoria-api/domain/cerrors"
)

type Medium struct {
	ID          string
	UserID      string
	UserSpaceID string
	Name        string
	Extension   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewMediumDTO struct {
	ID          string
	UserID      string
	UserSpaceID string
	Name        string
	Extension   string
}

func NewMedium(dto NewMediumDTO) *Medium {
	return &Medium{
		ID:          dto.ID,
		UserID:      dto.UserID,
		UserSpaceID: dto.UserSpaceID,
		Name:        dto.Name,
		Extension:   dto.Extension,
	}
}

var (
	ValidImageExtension = []string{"jpeg", "jpg", "png", "heic", "webp"}
	ValidVideoExtension = []string{"mp4", "mov", "3gp", "mkv", "ts", "webm"}
)

func (m Medium) GetNormalizedExtension() string {
	return strings.ToLower(m.Extension[1:])
}

func (m Medium) IsImage() bool {
	ownExt := m.GetNormalizedExtension()
	for _, ext := range ValidImageExtension {
		if ext == ownExt {
			return true
		}
	}
	return false
}

func (m Medium) IsVideo() bool {
	ownExt := m.GetNormalizedExtension()
	for _, ext := range ValidVideoExtension {
		if ext == ownExt {
			return true
		}
	}
	return false
}

func (m Medium) GetType() string {
	if m.IsImage() {
		return "image"
	}
	if m.IsVideo() {
		return "video"
	}

	panic(cerrors.NewNotImplemented(m.Extension))
}

func (m Medium) GetFileDirectoryPath() string {
	return "media/" + m.ID
}

func (m Medium) GetOriginalURL() string {
	return config.S3BucketHost + "/media/" + m.ID + "/original" + m.Extension
}

func (m Medium) GetTn240URL() string {
	return config.S3BucketHost + "/media/" + m.ID + "/240.png"
}
