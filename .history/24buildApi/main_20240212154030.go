package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Swift",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Sahil Dhargave",
			Website:  "http://bymetea",
		},
	})

	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "JS",
		CoursePrice: 29,
		Author: &Author{
			Fullname: "Sahil Dhargave",
			Website:  "http://buymecoffee",
		},
	})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":3002", r))

}

//controllers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by BuyMeCoffee"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)
	//fmt.Printf("Type of params: %v\n", params)
	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No courses found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if: boddy is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // &course -> address of variable / reference
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside JSON")
		return
	}

	// check for duplicate ID's
		params := mux.Vars(r)
	for _, course := range courses{
		if course.CourseName == params["name"] {
           json.NewEncoder(w).Encode("Please try Other title")
		   return
		}
	}
	// Generate unique id, string and
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create on Course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params := mux.Vars(r)

	// loop, id , remove from, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
		if course.CourseId != params["id"] {
			json.NewEncoder(w).Encode("NotFound")
			return
		}

	}

	// TODO: send a Response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//loop,id,remove(index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("We are successfully Delete Course")
			break
		}
	}

}
