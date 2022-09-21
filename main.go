package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

// var cancel context.CancelFunc

var database = "test"

// Connect to the database
func Connect() error {
	var err error

	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	ctx = context.Background()

	connectionString := fmt.Sprintf("mongodb+srv://%v:%v@%v",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("URL"),
	)

	fmt.Printf("connectionString: %v\n", connectionString)

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(
		connectionString,
	))

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return err
	}

	return nil
}

func Close() {
	//defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
}
