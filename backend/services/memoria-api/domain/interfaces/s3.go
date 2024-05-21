package interfaces

import "io"

type S3Client interface {
	GetPresignedPutObjectURL(dto S3ClientGetPresignedPutObjectURLDTO) (url string, err error)
	PutObject(dto S3ClientPutObjectDTO) (err error)
	DeleteFolder(dto S3ClientDeleteFolderDTO) (err error)
}

type S3ClientGetPresignedPutObjectURLDTO struct {
	Key string
}

type S3ClientPutObjectDTO struct {
	Key  string
	Body io.Reader
}

type S3ClientDeleteFolderDTO struct {
	Prefix string
}
