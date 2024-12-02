package awssdk

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	apigwmgmt "github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

const (
	Region          = "us-east-1"
	AccessKeyID     = "test"
	SecretAccessKey = "test"
	Endpoint        = "https://109e43c7.execute-api.localhost.localstack.cloud:4566/local"
)

type AWSSDKClient struct {
	sess  *session.Session
	apiGW *apigwmgmt.ApiGatewayManagementApi
}

func (a *AWSSDKClient) setup() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
		Endpoint:    aws.String(Endpoint),
	})
	if err != nil {
		log.Println(err)
		return
	}
	a.sess = sess
	a.apiGW = apigwmgmt.New(a.sess)
}

func New() *AWSSDKClient {
	i := &AWSSDKClient{}
	i.setup()

	return i
}

func (a *AWSSDKClient) PostToConnection(connID string, data any) error {
	jsonData, _ := json.Marshal(data)
	input := &apigwmgmt.PostToConnectionInput{
		ConnectionId: aws.String(connID),
		Data:         jsonData,
	}
	_, err := a.apiGW.PostToConnection(input)
	if err != nil {
		log.Println(err)
	}
	return err
}
