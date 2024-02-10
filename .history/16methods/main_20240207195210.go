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

	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of the user is:", u.Email)
}
