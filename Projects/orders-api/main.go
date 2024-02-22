package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
	
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
	
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.path))
	})

	http.HandlerFunc("/hi", func(w https.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hi")
	})

	log.Fatal(http.ListenAndServe(":8080",nil))
}	fmt.Fprintf(w,"Hello, %q",html.EscapeString(r.URL.path))
	})

	http.HandlerFunc("/hi", func(w https.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hi")
	})

	log.Fatal(http.ListenAndServe(":8080",nil))
}