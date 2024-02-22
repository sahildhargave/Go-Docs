package controllers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	// ... (unchanged)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	// ... (unchanged)
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

func getAllTasks() []primitive.M {
	// ... (unchanged)
}

func taskComplete(task string) {
	// ... (unchanged)
}

func insertOneTask(task models.ToDoList) {
	// ... (unchanged)
}

func undoTask(task string) {
	// ... (unchanged)
}

func deleteOneTask(task string) {
	// ... (unchanged)
}

func deleteAllTasks() int64 {
	// ... (unchanged)
}
