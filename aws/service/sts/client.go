package sts

import (
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sts"
)

func CreateSTSClient(session client.ConfigProvider) *sts.STS {
	return sts.New(session)
}
