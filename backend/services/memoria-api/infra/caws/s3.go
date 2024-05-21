package caws

import (
	"context"
	"time"

	"memoria-api/config"
	"memoria-api/domain/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Client struct {
	client *s3.Client
}

func NewS3Client(cfg aws.Config) (c interfaces.S3Client) {
	client := s3.NewFromConfig(cfg)
	c = &S3Client{
		client: client,
	}
	return
}

func (s *S3Client) GetPresignedPutObjectURL(dto interfaces.S3ClientGetPresignedPutObjectURLDTO) (presignedURL string, err error) {
	presignClient := s3.NewPresignClient(s.client)
	url, err := presignClient.PresignPutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(config.S3BucketName),
			Key:    aws.String(dto.Key),
		},
		s3.WithPresignExpires(time.Minute*15),
	)
	if err != nil {
		return
	}

	presignedURL = url.URL
	return
}

func (s *S3Client) PutObject(dto interfaces.S3ClientPutObjectDTO) (err error) {
	_, err = s.client.PutObject(
		context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(config.S3BucketName),
			Key:    aws.String(dto.Key),
			Body:   dto.Body,
		},
	)
	if err != nil {
		return
	}

	return
}

func (s *S3Client) DeleteFolder(dto interfaces.S3ClientDeleteFolderDTO) (err error) {
	ctx := context.TODO()

	listRes, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(config.S3BucketName),
		Prefix: aws.String(dto.Prefix),
	})
	if err != nil {
		return
	}

	// deleting the files under folder
	fileIdentifiers := make([]types.ObjectIdentifier, 0, len(listRes.Contents))
	for _, c := range listRes.Contents {
		fileIdentifiers = append(fileIdentifiers, types.ObjectIdentifier{Key: c.Key})
	}
	_, err = s.client.DeleteObjects(
		ctx,
		&s3.DeleteObjectsInput{
			Bucket: aws.String(config.S3BucketName),
			Delete: &types.Delete{
				Objects: fileIdentifiers,
			},
		},
	)

	// deleting the folder
	_, err = s.client.DeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(config.S3BucketName),
			Key:    aws.String(dto.Prefix),
		},
	)

	return
}
