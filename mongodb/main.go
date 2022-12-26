package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

// var cancel context.CancelFunc
var database = "contilab"

// Connect to the database
func init() {
	var err error

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	ctx = context.Background()

	// connection string to localhost
	connectionString := "mongodb://localhost:27017"

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		connectionString,
	))

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

}

func Close() {
	//defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}
