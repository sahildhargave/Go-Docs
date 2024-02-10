//

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome info of struct")

	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", "true", 23}

	fmt.Println("Info of User", hitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}