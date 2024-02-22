//☮✝☪🕉☸🕉☪✝

package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadEnv()
	createDBInstance()

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	connectingString := os.Getenv("MONGO_URL")
	mongo_db_name := os.Getenv("MONGO_DB_NAME")
	collName := os.Getenv("MONGO_DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectingString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")
	collection = client.Database(mongo_db_name).Collection(collName)
	fmt.Println("Collection instance created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllowTasks()
	json.NewEncoder(w).Encode(payload)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	// ... (unchanged)
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	// ... (unchanged)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// ... (unchanged)
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	// ... (unchanged)
}

//func getAllTasks() []primitive.M {
//	// ... (unchanged)
//}

func taskComplete(task string) {
	// ... (unchanged)
}

//func insertOneTask(task models.ToDoList) {
//	return
//}

func undoTask(task string) {
	// ... (unchanged)
}

func deleteOneTask(task string) {
	// ... (unchanged)
}

//func deleteAllTasks() int64 {
//	return
//}
