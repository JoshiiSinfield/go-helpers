package sts

import (
	errorHelper "github.com/JoshiiSinfield/go-helpers/aws/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetAccountNumberFromSession(session client.ConfigProvider) string {
	stsSvc := CreateSTSClient(session)
	input := &sts.GetCallerIdentityInput{}
	result, err := stsSvc.GetCallerIdentity(input)
	if err != nil {
		errorHelper.CastAWSError(err)
	}

	return aws.StringValue(result.Account)
}
