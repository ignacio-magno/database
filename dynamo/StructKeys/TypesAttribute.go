package StructKeys

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TypesAttribute[T any] struct {
	values    []interface{}
	keysQuery *KeysQuery[T]
}

func NewTypesAttribute[T any](values []interface{}, keysQuery *KeysQuery[T]) *TypesAttribute[T] {
	return &TypesAttribute[T]{
		values:    values,
		keysQuery: keysQuery,
	}
}

func (s *TypesAttribute[T]) GetTypesAttributes() []types.AttributeValue {
	attributeValues := make([]types.AttributeValue, len(s.values))

	for i, key := range s.values {
		marshalMap, err := attributevalue.Marshal(key)
		if err != nil {
			panic(err)
		}
		attributeValues[i] = marshalMap
	}

	return attributeValues
}

func (s *TypesAttribute[T]) GetKeyConditions() (map[string]types.Condition, error) {
	attributeValues := s.GetTypesAttributes()

	conditions := make(map[string]types.Condition, len(s.values))

	conditions[s.keysQuery.getPrimaryKey()] = types.Condition{
		ComparisonOperator: types.ComparisonOperatorEq,
		AttributeValueList: []types.AttributeValue{
			attributeValues[0],
		},
	}

	if s.keysQuery.haveTwoKeys && len(s.values) == 2 {
		conditions[s.keysQuery.getSecondaryKey()] = types.Condition{
			ComparisonOperator: types.ComparisonOperatorEq,
			AttributeValueList: []types.AttributeValue{
				attributeValues[1],
			},
		}
	}

	return conditions, nil
}

// GetAttributeWithKeys return the attribute value bind with the key in the model, is bound by index
func (s *TypesAttribute[T]) GetAttributeWithKeys() map[string]types.AttributeValue {
	attributeValues := make([]types.AttributeValue, len(s.values))

	for i, key := range s.values {
		marshalMap, err := attributevalue.Marshal(key)
		if err != nil {
			panic(err)
		}
		attributeValues[i] = marshalMap
	}

	conditions := make(map[string]types.AttributeValue, len(s.values))

	conditions[s.keysQuery.getPrimaryKey()] = attributeValues[0]

	if s.keysQuery.haveTwoKeys && len(s.values) == 2 {
		conditions[s.keysQuery.getSecondaryKey()] = attributeValues[1]
	}

	return conditions
}
