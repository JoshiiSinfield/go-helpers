package s3

import "github.com/aws/aws-lambda-go/events"

func ParseS3EventRecord(record events.S3EventRecord) (sourceBucket, objectKey string) {
	s3Record := record.S3

	// Retrieve the Bucket Name where the event originated from.
	sourceBucket = s3Record.Bucket.Name

	// Retrieve object key
	objectKey = s3Record.Object.Key

	return sourceBucket, objectKey
}
