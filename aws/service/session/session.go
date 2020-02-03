package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

func CreateSessionFromProfile(profileName, region string) client.ConfigProvider {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(region),
		},
		Profile:           profileName,
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func CreateSimpleSession(region string) client.ConfigProvider {
	mySess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Panicf("Unable to establish new Session, %v\n", err)
	}
	return mySess
}
