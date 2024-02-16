package main

import "fmt"

// Model for course -file
type Course struct {
 CourseId string `json:"courseId"`
 CourseName string `json:"courseName"`
 Description string `json:"description"`
 CoursePrice int `json:"price"`
 
}

type Author struct {
  Fullname string `json:"fullname"`
  Website string  `json:"website"`

}
func main(){
	fmt.Println("API Functionality");

}