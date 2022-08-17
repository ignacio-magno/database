package database

import "go.mongodb.org/mongo-driver/mongo"

type Delete struct {
	Db
}

func (db *Delete) Delete() (*mongo.DeleteResult, error) {
	res, err := client.Database(database).Collection(db.getNameCollection()).DeleteMany(ctx, db.getFilter())
	if err != nil {
		return res, err
	}
	return res, nil
}

func (db *Delete) DeleteOne() (*mongo.DeleteResult, error) {
	res, err := client.Database(database).Collection(db.getNameCollection()).DeleteOne(ctx, db.getFilter())
	if err != nil {
		return res, err
	}
	return res, nil
}
