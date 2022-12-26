package mongodb_test

import (
	"fmt"
	"github.com/ignacioMagno/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type Test struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

var r = mongodb.NewRepository[Test]("test", "test-package")

func TestFind(t *testing.T) {
	result := r.Find(bson.M{r.Field(1): "test nachita"})

	for _, test := range result {
		fmt.Println(test)
	}
}

func TestSave(t *testing.T) {
	r.Save(Test{Name: "test nachita"})
}

func TestById(t *testing.T) {
	result := r.FindById("63a901fefbd11b63c6a0bb8a")
	fmt.Println(result)
}
