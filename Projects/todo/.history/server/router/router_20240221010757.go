package router

import (
	"github.com/gorilla/mux"
	"github.com/sahildhargave/todo/contros"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", controllers.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task", controllers.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task", controllers.undoTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task", controllers.DeleteTask).Methods("DELETE", "OPTIONS")
	return router

}
