package router

import (
	"github.com/gorilla/mux"
	"github.com/sahildhargave/todo/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middle.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middle.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task", middle.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task", middle.undoTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task", middle.DeleteTask).Methods("DELETE", "OPTIONS")
	return router

}
