package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"reflect"
	"strings"
)

// StructKeys take the T type and extract the primary and secondary keys
type StructKeys[T any] struct {
}

func (s *StructKeys[T]) GetPrimaryKey() string {
	var t T
	return strings.Split(reflect.TypeOf(t).Field(0).Tag.Get("dynamodbav"), ",")[0]
}

func (s *StructKeys[T]) GetSecondaryKey() string {
	var t T
	return strings.Split(reflect.TypeOf(t).Field(1).Tag.Get("dynamodbav"), ",")[0]
}

func (s *StructKeys[T]) GetKeyConditions(keys []interface{}) map[string]types.Condition {
	if len(keys) == 0 {
		panic("keys is required for query")
	}
	attributeValues := make([]types.AttributeValue, len(keys))

	for i, key := range keys {
		marshalMap, err := attributevalue.Marshal(key)
		if err != nil {
			panic(err)
		}
		attributeValues[i] = marshalMap
	}

	conditions := make(map[string]types.Condition, len(keys))

	conditions[s.GetPrimaryKey()] = types.Condition{
		ComparisonOperator: types.ComparisonOperatorEq,
		AttributeValueList: []types.AttributeValue{
			attributeValues[0],
		},
	}

	if len(keys) > 1 {
		conditions[s.GetSecondaryKey()] = types.Condition{
			ComparisonOperator: types.ComparisonOperatorEq,
			AttributeValueList: []types.AttributeValue{
				attributeValues[1],
			},
		}
	}

	if len(keys) > 2 {
		panic("only two keys max are allowed")
	}

	return conditions
}

func (s *StructKeys[T]) GetAttributeValue(keys []interface{}) map[string]types.AttributeValue {
	if len(keys) == 0 {
		panic("keys is required for query")
	}
	attributeValues := make([]types.AttributeValue, len(keys))

	for i, key := range keys {
		marshalMap, err := attributevalue.Marshal(key)
		if err != nil {
			panic(err)
		}
		attributeValues[i] = marshalMap
	}

	conditions := make(map[string]types.AttributeValue, len(keys))

	conditions[s.GetPrimaryKey()] = attributeValues[0]

	if len(keys) > 1 {
		conditions[s.GetSecondaryKey()] = attributeValues[1]
	}

	if len(keys) > 2 {
		panic("only two keys max are allowed")
	}

	return conditions
}
