package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int      `json:"Price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"Tags"`
}

func main() {
	fmt.Println("Information about the JSON application and how to deal with")
	EncodeJSON()
}

func EncodeJSON() {
	BuyMeCoffee := []course{
		{
			Name:     "Web Development Fundamentals",
			Price:    50,
			Platform: "codingacademy.com",
			Password: "abc123",
			Tags:     []string{"web-dev", "html", "css"},
		},
		{
			Name:     "Advanced JavaScript",
			Price:    75,
			Platform: "javascriptmasters.net",
			Password: "xyz789",
			Tags:     []string{"web-dev", "js", "es6"},
		},
		{
			Name:     "Data Science with Python",
			Price:    120,
			Platform: "datasciencelab.org",
			Password: "pandas456",
			Tags:     []string{"data-science", "python", "pandas"},
		}, {
			Name:     "Angular",
			Price:    100,
			Platform: "angular.org",
			Password: "angular456",
			Tags:     nil,
		},
	}

	//package this data as JSON data

	ourJson, err := json.MarshalIndent(BuyMeCoffee, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", ourJson)
	//	fmt.Println(string(ourJson)) //print out the JSON data

}
