package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_CreateGetObjectParams(t *testing.T) {
	type args struct {
		bucket    string
		objectKey string
	}
	tests := []struct {
		name string
		args args
		want *s3.GetObjectInput
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				bucket:    "bucket1",
				objectKey: "path/to/object1",
			},
			want: &s3.GetObjectInput{
				Bucket: aws.String("bucket1"),
				Key:    aws.String("path/to/object1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGetObjectParams(tt.args.bucket, tt.args.objectKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGetObjectParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateCopyObjectParams(t *testing.T) {
	type args struct {
		destBucket   string
		sourceBucket string
		objectKey    string
	}
	tests := []struct {
		name string
		args args
		want *s3.CopyObjectInput
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				destBucket:   "s3-transfer-test-dest",
				sourceBucket: "s3-transfer-test-source",
				objectKey:    "test/object/key/test.txt",
			},
			want: &s3.CopyObjectInput{
				Bucket:     aws.String("s3-transfer-test-dest"),
				CopySource: aws.String(filepath.Join("s3-transfer-test-source", "test/object/key/test.txt")),
				Key:        aws.String("test/object/key/test.txt"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCopyObjectParams(tt.args.destBucket, tt.args.sourceBucket, tt.args.objectKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCopyObjectParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateGetObjectParams(t *testing.T) {
	type args struct {
		bucket    string
		objectKey string
	}
	tests := []struct {
		name string
		args args
		want *s3.GetObjectInput
	}{
		// TODO: Add test cases.
		{
			name: "ShouldReturnACorrectlyFilledGetObjectInput",
			args: args{
				bucket:    "test-getinput-bucket-123456",
				objectKey: "test/get/object/input",
			},
			want: &s3.GetObjectInput{
				Key:    aws.String("test/get/object/input"),
				Bucket: aws.String("test-getinput-bucket-123456"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGetObjectParams(tt.args.bucket, tt.args.objectKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGetObjectParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
