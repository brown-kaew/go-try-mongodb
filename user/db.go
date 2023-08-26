package user

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TryAddAndFindData() (*mongo.Collection, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))

	collection := client.Database("mydb").Collection("user")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx,
		bson.D{
			{Key: "name", Value: "pi"},
			{Key: "value", Value: 3.14159},
		})
	id := res.InsertedID
	log.Printf("%v\n", id)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		log.Printf("%v\n", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return collection, func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
}
