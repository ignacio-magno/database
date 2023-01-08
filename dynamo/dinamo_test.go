package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"testing"
)

type TestStruct struct {
	IdBusiness string `dynamodbav:"var_id_client"`
	TypeEntity string `dynamodbav:"var_id_business_type_entity"`
	Data       string `dynamodbav:"data"`
}

var repo = NewRepositoryDynamo[TestStruct]("contilab", true)

func TestMain(m *testing.M) {
	m.Run()
}

func TestSave(t *testing.T) {
	err := repo.SaveOrReplace(TestStruct{
		IdBusiness: "nacho",
		TypeEntity: "1",
		Data:       "data",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	res, err := repo.Find([]interface{}{"nacho", "1"})
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestFindOne(t *testing.T) {
	res, err := repo.FindOne([]interface{}{"nacho", "1"})
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestUpdate(t *testing.T) {
	j, err := repo.Update([]interface{}{"nacho", "1"}, map[string]types.AttributeValueUpdate{
		"data": {
			Action: types.AttributeActionPut,
			Value:  &types.AttributeValueMemberS{Value: "data2"},
		},
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(j)
}

func TestDelete(t *testing.T) {
	err := repo.Delete([]interface{}{"nacho", "1"})
	if err != nil {
		t.Error(err)
	}
}
