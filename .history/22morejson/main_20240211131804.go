
package main

import "fmt"


type course struct{
	Name string
	Price int
	platform string
	Password string 
	Tags []string
}

func main(){
	fmt.Println("Information about the JSON application and how to deal with");
    
}

func EncodeJson(){
	buymecoffee := []course{
		"Go Bootcamp",
		90,
		"sahil.in",
		"abe345",
		[]string{
			"web-dev",
			"js"
		},
	}
}