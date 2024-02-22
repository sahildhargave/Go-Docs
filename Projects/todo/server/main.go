package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sahildhargave/todo/router"
)

func main() {
	r := router.Router()
	port := ":5000"

	fmt.Printf("Server is running on http://localhost%s 🚀\n", port)

	log.Fatal(http.ListenAndServe(port, r))
}
