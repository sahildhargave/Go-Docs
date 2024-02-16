package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"Price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` // omit empty slices
}

func main() {
	fmt.Println("Information about the JSON application and how to deal with")
	EncodeJSON()
	fmt.Println("😁😂🤣😃😃😄🥰😘😍😎😍😘🥰")
	decodeJson()
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

func decodeJson() {
	BuyMeCoffee := []byte(`
	{
                "coursename": "Web Development Fundamentals",        
                "Price": 50,
                "website": "codingacademy.com",
                "tags": ["web-dev","html","css"]
        }

	`)

	var ourcourse course
	checkValid := json.Valid(BuyMeCoffee) //check if valid JSON
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(BuyMeCoffee, &ourcourse) //decode the JSON into GoLang
		fmt.Printf("%#v\n", ourcourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(BuyMeCoffee, &myOnlineData)
	fmt.P
}
