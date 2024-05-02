package caws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const region = "ap-northeast-1"

// LoadConfig assumes following environment variables are set
// - AWS_ACCESS_KEY_ID
// - AWS_SECRET_ACCESS_KEY
func LoadConfig(ctx context.Context) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, config.WithRegion(region))
}
