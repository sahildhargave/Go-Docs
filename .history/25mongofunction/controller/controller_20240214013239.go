package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/sahildhargave/25mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sahildhargave5253:1KJk5mrwbJeGlq0d@cluster0.gpxr5wx.mongodb.net/?retryWrites=true&w=majority"

const dbName = "Netflix"
const colName = "watchlist"

// most important'
var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB conntection success")
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

//MONGODB helpers -file

// insert one record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.TODO(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

// update 1 record
//should by hard learn and keep in mind
//TODO - Difference between bson.m and bson.D
// from stackoverflow

func updateOnMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	//😀😁 THERE ARE TWO BSON.M and BSON.D
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Modified %v documents\n", result.ModifiedCount)
}


func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	collection.D
}