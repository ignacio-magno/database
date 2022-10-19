package dynamo

type TestStruct struct {
	IdBusiness string `dynamodbav:"id-business"`
	TypeEntity int    `dynamodbav:"type-entity"`
	PagaIva    bool   `dynamodbav:"have-iva"`
}

type RepositoryTest struct {
	Repository[TestStruct]
}

func NewRepositoryTest() *RepositoryTest {
	return &RepositoryTest{
		Repository[TestStruct]{
			NameCollection: "contilab-password",
		},
	}
}
