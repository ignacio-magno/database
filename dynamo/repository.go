package dynamo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	StructKeys "github.com/ignacio-magno/database/dynamo/StructKeys"
)

type Repository[T any] struct {
	*StructKeys.KeysQuery[T]
	NameCollection string
}

func NewRepositoryDynamo[T any](nameDatabase string, haveTwoKeys bool) *Repository[T] {
	return &Repository[T]{
		NameCollection: nameDatabase,
		KeysQuery:      StructKeys.NewKeysQuery[T](haveTwoKeys),
	}
}

func (m *Repository[T]) GetNameCollection() string {
	return m.NameCollection
}

func (m *Repository[T]) GenerateProjectionExpressionExclude(filters []string) *string {
	var t T
	return GenerateProjectionExpressionExclude(t, filters)
}

// Find return array of elements founded
// @param keys
// * required
// primary key = keys[0]
// * optional
// secondary key = keys[1]
func (m *Repository[T]) Find(keys []interface{}, queryInputHandler ...func(input *dynamodb.QueryInput)) ([]T, error) {
	var (
		queryInput dynamodb.QueryInput
		err        error
	)

	bta, err := m.BuildTypesAttribute(keys, false)
	if err != nil {
		return nil, err
	}

	queryInput.TableName = aws.String(m.GetNameCollection())
	queryInput.KeyConditions, err = bta.GetKeyConditions()
	if err != nil {
		return nil, err
	}

	for _, f := range queryInputHandler {
		f(&queryInput)
	}

	query, err := DynamoClient.Query(context.Background(), &queryInput)

	if err != nil {
		return nil, err
	}

	// marshal response
	result := make([]T, len(query.Items))

	for i, item := range query.Items {
		var data T
		err = attributevalue.UnmarshalMap(item, &data)
		if err != nil {
			return nil, err
		}
		result[i] = data
	}

	return result, err
}

// FindOne return one element founded, if not found return error
func (m *Repository[T]) FindOne(id []interface{}) (T, error) {
	find, err := m.Find(id)
	var t T
	if err != nil {
		return t, err
	}

	if len(find) == 0 {
		return t, fmt.Errorf("not found")
	}

	if len(find) > 1 {
		return t, fmt.Errorf("more than one result")
	}

	return find[0], nil
}

// Update if the table have 2 keys, then set boot keys
func (m *Repository[T]) Update(keys []interface{}, update map[string]types.AttributeValueUpdate) (T, error) {
	var (
		t   T
		err error
	)

	bta, err := m.BuildTypesAttribute(keys, true)
	if err != nil {
		return t, err
	}

	res, err := DynamoClient.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName:        aws.String(m.GetNameCollection()),
		Key:              bta.GetAttributeWithKeys(),
		AttributeUpdates: update,
	})

	if err != nil {
		return t, err
	}

	var data T
	err = attributevalue.UnmarshalMap(res.Attributes, &data)

	return data, err
}

// SaveOrReplace save a new element or replace if exist previous
func (m *Repository[T]) SaveOrReplace(document T) error {

	marshalMap, err := attributevalue.MarshalMap(document)
	if err != nil {
		return err
	}

	item, err := DynamoClient.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(m.GetNameCollection()),
		Item:      marshalMap,
	})
	_ = item

	return err
}

func (m *Repository[T]) SaveMany(documents []T) error {
	for _, val := range documents {
		err := m.SaveOrReplace(val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Repository[T]) Delete(keys []interface{}) error {
	att, err := m.BuildTypesAttribute(keys, true)
	if err != nil {
		return err
	}

	_, err = DynamoClient.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(m.GetNameCollection()),
		Key:       att.GetAttributeWithKeys(),
	})

	return err
}
