package main

import "fmt"

// Model for course -file
type Course struct {
 courseId string `json:"courseId"`
}

type Author struct {
  Fullname string `json:"fullname"`
  Website string  `json:"website"`

}
func main(){
	fmt.Println("API Functionality");

}