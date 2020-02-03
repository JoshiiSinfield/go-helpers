package errors

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func CastAWSError(err error) {
	// cast err to awserr.Error to get the Code and Message from an error.
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case s3.ErrCodeObjectNotInActiveTierError:
			log.Println(s3.ErrCodeObjectNotInActiveTierError, aerr.Error())
		case s3.ErrCodeNoSuchKey:
			log.Println(s3.ErrCodeNoSuchKey, aerr.Error())
		default:
			log.Println(aerr.Error())
		}
	} else {
		log.Println(err.Error())
	}
	return
}
