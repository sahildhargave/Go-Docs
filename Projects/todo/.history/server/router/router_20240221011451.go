package router

import (
	"github.com/gorilla/mux"
	"github.com/sahildhargave/todo/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task", middleware.undoTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	return router

}
