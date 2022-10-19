package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Update struct {
	query ISaveOrModify
}

func (db *Update) UpdateOne() (*mongo.UpdateResult, error) {
	res, err := client.Database(database).Collection(db.query.getNameCollection()).UpdateOne(ctx, db.query.getFilter(), db.query.getData())
	if err != nil {
		return res, err
	}
	return res, nil
}

func (db *Update) Update() (*mongo.UpdateResult, error) {
	res, err := client.Database(database).Collection(db.query.getNameCollection()).UpdateMany(ctx, db.query.getFilter(), db.query.getData())
	if err != nil {
		return res, err
	}
	return res, nil
}
