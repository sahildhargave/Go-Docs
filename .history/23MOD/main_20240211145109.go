package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("MOD in golang")
}

func greeter() {
	fmt.Println("Hey there I am Joscript")

}

func serverHome(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte(""))
}
