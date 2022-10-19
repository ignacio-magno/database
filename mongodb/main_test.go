package mongodb_test

import (
	"backend-contilab/database/mongodb"
	"fmt"
	"testing"

	database "backend-contilab/database"
	"go.mongodb.org/mongo-driver/bson"
)

type testStruct struct {
	Test string `bson:"test" json:"test"`
}

func TestSaveOne(m *testing.T) {
	defer mongodb.Close()

	s := mongodb.Save{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
		},
		Add: mongodb.Add{
			Data: testStruct{
				Test: "test",
			},
		},
	}

	i, err := s.SaveOne()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("i: %v\n", i.InsertedID)
}

func TestSaveMany(t *testing.T) {
	defer mongodb.Close()

	s := mongodb.Save{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
		},
		Add: mongodb.Add{
			Datas: []interface{}{
				testStruct{
					Test: "test",
				},
				testStruct{
					Test: "test",
				},
			},
		},
	}

	id, err := s.Save()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	for i := range id.InsertedIDs {
		fmt.Printf("id: %v\n", id.InsertedIDs[i])
	}
}

func TestFind(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	var tests []testStruct
	find := mongodb.Find{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Bind: &tests,
	}

	err = find.Find()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	for i, ts := range tests {
		fmt.Printf("%v: %v\n", i, ts.Test)
	}
}

func TestFindOne(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	var testObj testStruct
	find := mongodb.Find{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Bind: &testObj,
	}

	err = find.FindOne()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("%v\n", testObj.Test)
}

func TestFindOneAndDelete(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	var testObj testStruct

	find := mongodb.FindAndModify{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Bind: &testObj,
	}
	err = find.FindOneAndDelete()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}

func TestFindOneAndReplace(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	var testObj testStruct
	testObj.Test = "test2"
	find := mongodb.FindAndModify{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Data: testObj,
		Bind: &testObj,
	}

	err = find.FindOneAndReplace()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("testObj: %v\n", testObj)
}

func TestFindOneAndUpdate(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	var testObject testStruct
	testObject.Test = "test2"
	find := mongodb.FindAndModify{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Data: bson.D{{"$set", bson.E{Value: "ignacio"}}},
		Bind: &testObject,
	}

	err = find.FindOneAndUpdate()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("testObject: %v\n", testObject)
}

func TestDeleteOne(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	deleteQuery := mongodb.Delete{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
	}

	_, err = deleteQuery.DeleteOne()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}

func TestDelete(t *testing.T) {
	err := database.Connect()
	if err != nil {
		return
	}
	defer mongodb.Close()

	deleteQuery := mongodb.Delete{
		Db: mongodb.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
	}

	_, err = deleteQuery.Delete()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
