package database_test

import (
	"fmt"
	"testing"

	"github.com/ignacioMagno/database"
	"go.mongodb.org/mongo-driver/bson"
)

type testStruct struct {
	Test string `bson:"test" json:"test"`
}

func TestSaveOne(m *testing.T) {
	s := database.Save{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
		},
		Add: database.Add{
			Data: testStruct{
				Test: "test",
			},
		},
	}

	_, err := s.SaveOne()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}

func TestSaveMany(t *testing.T) {
	s := database.Save{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
		},
		Add: database.Add{
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

	_, err := s.Save()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}

func TestFind(t *testing.T) {
	var tstruct []testStruct
	find := database.Find{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Bind: &tstruct,
	}

	err := find.Find()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	for i, ts := range tstruct {
		fmt.Printf("%v: %v\n", i, ts.Test)
	}
}

func TestFindOne(t *testing.T) {
	var tstruct testStruct
	find := database.Find{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Bind: &tstruct,
	}

	err := find.FindOne()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	fmt.Printf("%v\n", tstruct.Test)
}

func TestFindOneAndDelete(t *testing.T) {
	var tstruct testStruct
	find := database.FindAndModify{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Bind: &tstruct,
	}
	find.FindOneAndDelete()
}

func TestFindOneAndReplace(t *testing.T) {
	var tstruct testStruct
	tstruct.Test = "test2"
	find := database.FindAndModify{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Data: tstruct,
		Bind: &tstruct,
	}

	err := find.FindOneAndReplace()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	fmt.Printf("tstruct: %v\n", tstruct)
}

func TestFindOneAndUpdate(t *testing.T) {
	var tstruct testStruct
	tstruct.Test = "test2"
	find := database.FindAndModify{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
		Data: bson.D{{"$set", bson.E{Value: "ignacio"}}},
		Bind: &tstruct,
	}

	err := find.FindOneAndUpdate()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	fmt.Printf("tstruct: %v\n", tstruct)
}

func TestDeleteOne(t *testing.T) {
	delete := database.Delete{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
	}

	_, err := delete.DeleteOne()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}

func TestDelete(t *testing.T) {
	delete := database.Delete{
		Db: database.Db{
			NameCollection: "test",
			Database:       "test",
			Filter:         bson.D{{Key: "test", Value: "test"}},
		},
	}

	_, err := delete.Delete()
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
