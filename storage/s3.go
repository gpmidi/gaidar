package storage

import "github.com/aws/aws-sdk-go/service/s3"

type ImageStore struct {
	BucketName string
	Bucket     s3.Bucket
}


