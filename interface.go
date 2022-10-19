package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDatabaseRepository[T any] interface {
	Find(id interface{}) ([]T, error)
	FindOne(id interface{}) (T, error)
	FindBy(filter bson.M) ([]T, error)
	Update(ids interface{}, document T) (T, error)
	Save(document T) (interface{}, error)
	SaveMany(documents []T) ([]interface{}, error)
	Delete(id string) error
	DeleteByFilter(filter bson.M) (*mongo.DeleteResult, error)
	FindAndReplace(data T, filter bson.M) (replaced T, err error)
	FindAndUpdate(data bson.M, filter bson.M) (updated T, err error)
}
