package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("MOD in golang")
	greeter()
	r := mux.NewRouter
}

func greeter() {
	fmt.Println("Hey there I am Joscript")

}

func serverHome(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("<h1>Welcome to buymecoffee on YT</h1>"))
}
