package main

import (
	"log"
	"net/http"
    
	"../pkg/routes/bookstore-routes.go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090",r))

}