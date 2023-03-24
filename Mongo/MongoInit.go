package mongoinit

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func MongoConnInit() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI("mongodb+srv://rahul:1234@cluster0.hzrgq2w.mongodb.net/?retryWrites=true&w=majority")
	options.SetMaxConnIdleTime(time.Second * 20)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		fmt.Println(err)
	}

	MongoDB = client.Database("Product")
	return MongoDB
}
