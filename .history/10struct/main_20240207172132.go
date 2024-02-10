//

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome info of struct")

	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 23}

	fmt.Println("Info of User", hitesh)
	fmt.Printf("\n hitesh info are: %+v\n", hitesh)

	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
