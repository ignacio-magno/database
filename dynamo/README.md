# Doc

build repository with struct model

```go
package test

// always set the primary key in the first field, and the sort key in the second field if exist
type TestStruct struct {
	IdBusiness string `dynamodbav:"var_id_client"`               //primary key
	TypeEntity string `dynamodbav:"var_id_business_type_entity"` // sort key tag, optional
	Data       string `dynamodbav:"data"`
}

// first parameter is table name
// second parameter is if have two keys
var repo = NewRepositoryDynamo[TestStruct]("contilab", true)

```