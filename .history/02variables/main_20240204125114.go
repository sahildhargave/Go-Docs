package main

import "fmt"

func main() {
	var username string = "Sahil"
	fmt.Println(username)
	fmt.Printf("variables is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type :%T\n", smallVal)

	var smallfloat float32 = 255.67854679535486785564
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type :%T\n", smallfloat)

	var largefloat float64 = 255.67854679535486785564
	fmt.Println(largefloat)
	fmt.Printf("Variable is of type :%T\n", largefloat)

	//
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type :%T\n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type :%T\n", anotherString)

}
