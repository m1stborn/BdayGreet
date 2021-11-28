package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	//"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

var MongoDB = MongoDatabase{}

func (d *MongoDatabase) Init(uri string) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	d.Client = client

	collection := client.Database("BdayGreet").Collection("user")
	d.Collection = collection

	// insert all the data
	var users []interface{}
	for _, user := range DBRecords {
		users = append(users, user)
	}
	//d.InsertMany(users)

}

func (d *MongoDatabase) InsertMany(toInsert []interface{}) {
	insertResult, err := d.Collection.InsertMany(context.TODO(), toInsert)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
}

func (d *MongoDatabase) FindByBday(month int, day int) ([]bson.M, error) {
	filter := bson.M{
		"$expr": bson.M{
			"$and": []bson.M{
				{"$eq": []interface{}{bson.M{"$month": "$birth"}, month}},
				{"$eq": []interface{}{bson.M{"$dayOfMonth": "$birth"}, day}},
			},
		},
	}

	filterCursor, err := d.Collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterCursor.All(context.TODO(), &usersFiltered); err != nil {
		log.Fatal(err)
	}

	return usersFiltered, nil
}
