
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
	BuyMeC := []course{
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
		},
	}


}