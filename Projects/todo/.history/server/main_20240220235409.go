package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sahil/todo/router"
)

func main(){
	r:= router.Router()
	fmt.Println("Staring on 👩‍💻👨‍💻👩‍💻👨‍💻")
	log.Fatal(http.ListenAndServe(":5000",r))
}