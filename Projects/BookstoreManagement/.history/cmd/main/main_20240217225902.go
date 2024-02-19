package main

import (
	"log"
	"net/http"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/mux"
	 "github.com"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))

}
