package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Save struct {
	Db
	Add
}

func (s *Save) SaveOne() (*mongo.InsertOneResult, error) {
	if s.getData() != nil {
		res, err := client.Database(database).Collection(s.getNameCollection()).InsertOne(ctx, s.getData())
		fmt.Printf("err: %v\n", err)
		fmt.Printf("res: %v\n", res)
		return res, err
	} else {
		return nil, fmt.Errorf("no data to save")
	}
}

func (s *Save) Save() (*mongo.InsertManyResult, error) {
	if s.getDatas() != nil {
		res, err := client.Database(database).Collection(s.getNameCollection()).InsertMany(ctx, s.getDatas())
		if err != nil {
			return res, err
		}
		return res, nil
	} else {
		return nil, fmt.Errorf("no array of data to save")
	}
}
