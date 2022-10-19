package dynamo

import (
	"fmt"
	"testing"
)

var repoTest = NewRepositoryTest()

func TestDynamo(t *testing.T) {

	find, err := repoTest.Find("123")
	if err != nil {
		return
	}

	for _, testStruct := range find {
		fmt.Println(testStruct)
	}

}
