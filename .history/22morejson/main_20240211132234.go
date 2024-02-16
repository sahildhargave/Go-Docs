
package main

import "fmt"


type course struct{
	Name string
	Price int
	Platform string
	Password string 
	Tags []string
}

func main(){
	fmt.Println("Information about the JSON application and how to deal with");
    
}

func EncodeJSON() {
	buymecoffee := course{
		{
        Name:     "Go Bootcamp",
		Price:    90,
		Platform: "sahil.in",
		Password: "abe345",
		Tags:     []string{"web-dev", "js"},
		},
			{
        Name:     "Go Bootcamp",
		Price:    90,
		Platform: "sahil.in",
		Password: "abe345",
		Tags:     []string{"web-dev", "js"},
		}
	}
	}


}