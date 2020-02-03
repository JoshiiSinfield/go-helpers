package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

func DownloadFromS3(client *s3.S3, bucket, key string) ([]byte, error) {
	getObjectParams := CreateGetObjectParams(bucket, key)

	result, err := client.GetObject(getObjectParams)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, result.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}