package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SecondCollection struct {
	Name string `bson:"name"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// https://www.mongodb.com/ja-jp/docs/drivers/go/current/usage-examples/insertOne/
	// こちらのコードを参考にして、InsertOneを利用
	coll := client.Database("my_database").Collection("second_collection")
	newSecondCollection := SecondCollection{Name: "8282"}
	result, err := coll.InsertOne(context.TODO(), newSecondCollection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("_id: %s\n", result.InsertedID)

	// second_collectionの取得
	var secondCollection SecondCollection
	err = coll.FindOne(context.TODO(), bson.D{{"name", newSecondCollection.Name}}).Decode(&secondCollection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name: %s\n", secondCollection.Name)

}
