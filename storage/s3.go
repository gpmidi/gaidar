package storage

import "github.com/aws/aws-sdk-go/service/s3"

// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/
import "github.com/aws/aws-sdk-go/service/s3"

type ImageStore struct {
	BucketName string
	Bucket     s3.Bucket
}


