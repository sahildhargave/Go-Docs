package controller

import (
	"context"

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
	client,err := mongo.Connect(context.TODO(), clientOption) 

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB conntection success")
	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}
