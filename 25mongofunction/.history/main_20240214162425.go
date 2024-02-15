package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sahildhargave/25mongodb/router"
)

func main() {
	fmt.Println("working with mongodb")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("Listening at port 4000 ...")

}
