package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynamoClient *dynamodb.Client

func init() {
	defaultConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	dynamoClient = dynamodb.NewFromConfig(defaultConfig)
}
