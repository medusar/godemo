package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
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

type User struct {
	Name string `bson:"name"`
	Date int64  `bson:"date"`
}

//{"$lt":121212121}
type DateCond struct {
	Date int64 `bson:"$lt"`
}

//{"date":{"$lt":121212121}}
type FindCond struct {
	DateCond *DateCond `bson:"date"`
}

func main() {
	client, err := Connect("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("my_db")
	collection := database.Collection("user")
	for i := 0; i < 10; i++ {
		_, err := collection.InsertOne(context.TODO(), &User{"Jason" + strconv.Itoa(i), time.Now().Unix()})
		if err != nil {
			log.Fatal(err)
		}
	}

	cur, err := collection.Find(context.TODO(), &FindCond{&DateCond{time.Now().Unix()}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		user := &User{}
		cur.Decode(user)
		log.Println(user)
	}

	many, err := collection.DeleteMany(context.TODO(), &FindCond{&DateCond{time.Now().Unix()}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("deletecd:", many)
}
