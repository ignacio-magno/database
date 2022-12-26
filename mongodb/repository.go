package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type Repository[T any] struct {
	collection string
	database   string
}

func NewRepository[T any](collection string, database string) *Repository[T] {
	return &Repository[T]{
		collection: collection,
		database:   database,
	}
}

func (r *Repository[T]) Find(filter interface{}) []T {
	fmt.Println(filter)
	find, err := client.Database(r.database).Collection(r.collection).Find(ctx, filter)
	if err != nil {
		return nil
	}

	if find.Err() != nil {
		panic(err.Error())
	}

	var result []T

	err = find.All(ctx, &result)
	if err != nil {
		return nil
	}

	return result
}

func (r *Repository[T]) Save(model T) T {
	_, err := client.Database(r.database).Collection(r.collection).InsertOne(ctx, model)
	if err != nil {
		panic(err.Error())
	}

	return model
}

func (r *Repository[T]) FindById(id string) T {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err.Error())
	}

	find := client.Database(r.database).Collection(r.collection).FindOne(ctx, bson.M{"_id": _id})

	if find.Err() != nil {
		panic(find.Err().Error())
	}

	var result T

	err = find.Decode(&result)
	if err != nil {
		panic(err.Error())
	}

	return result
}

func (r *Repository[T]) DeleteOne(filter interface{}) (deleteOk bool) {
	resp, err := client.Database(r.database).Collection(r.collection).DeleteOne(ctx, filter)
	if err != nil {
		panic(err.Error())
	}

	deleteOk = resp.DeletedCount > 0
	return
}

func (r *Repository[T]) DeleteMany(filter interface{}) (countDeleted int64) {
	resp, err := client.Database(r.database).Collection(r.collection).DeleteMany(ctx, filter)
	if err != nil {
		panic(err.Error())
	}

	countDeleted = resp.DeletedCount
	return
}

// Field reflect field from struct and return the name of field
func (r *Repository[T]) Field(index int) string {
	var t T
	return reflect.TypeOf(t).Field(index).Tag.Get("bson")
}
