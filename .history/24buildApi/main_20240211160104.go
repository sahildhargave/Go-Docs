package main

import (
	"fmt"
	"net/http"
)

// Model for course -file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	Description string  `json:"description"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake Database
var courses []Course

// middleware , helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.Description == "" && c.CoursePrice <= 0
}

func main() {
	fmt.Println("API Functionality")

}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by BuyMeCoffee"))
}


func getAllCourses(w http.ResponseWriter,r *http.Request){
   fmt.Println("Get all courses")
   w.Header().Set("Content-Type")
}