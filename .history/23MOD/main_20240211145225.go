package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("MOD in golang")
	greeter()
	r := mux.New
}

func greeter() {
	fmt.Println("Hey there I am Joscript")

}

func serverHome(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("<h1>Welcome to buymecoffee on YT</h1>"))
}
