package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"path/filepath"
)

func CreateGetObjectParams(bucket string, objectKey string) *s3.GetObjectInput {
	return &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}
}

func CreateCopyObjectParams(destBucket, sourceBucket, objectKey string) *s3.CopyObjectInput {
	return &s3.CopyObjectInput{
		Bucket:     aws.String(destBucket),
		CopySource: aws.String(filepath.Join(sourceBucket, objectKey)),
		Key:        aws.String(objectKey),
	}
}

func CreateHeadObjectParams(bucket, objectKey string) *s3.HeadObjectInput {
	return &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}
}

func CreatePutObjectParams(buffer []byte, bucket, filePath string) *s3.PutObjectInput {
	return &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader(buffer),
	}
}
