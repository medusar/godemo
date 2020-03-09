package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Connect(uri string, timeout time.Duration) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	client, err := Connect("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("my_db")
	collection := database.Collection("my_collection")

	ret, err := collection.InsertOne(context.TODO(), bson.M{"uid": 1001, "name": "jack"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted, id:", ret.InsertedID)

	find := options.Find()
	//find.SetSkip(0)
	//find.SetLimit(2)
	cur, err := collection.Find(context.TODO(), bson.M{"name": "jack"}, find)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := struct {
			Name string `bson:"name"`
			Uid  int32  `bson:"uid"`
		}{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...
		log.Println("result:", result)

		// To get the raw bson bytes use cursor.Current
		//raw := cur.Current
		// do something with raw...
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
