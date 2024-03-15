package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoClient *dynamodb.Client

func LoadDefaultClient() {
	defaultConfig, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	DynamoClient = dynamodb.NewFromConfig(defaultConfig)
}
