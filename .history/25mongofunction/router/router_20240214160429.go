//😁😂 Router

package router

import (
	"github.com/gorilla/mux"
	"github.com/sahildhargave/25mongodb/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.GetMyAllMovies).Methods("GET")
	return
}
