package main

import "fmt"

func main() {
	fmt.Println("Welcome to info of slice")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitList)

	fruitList = append(fruitList, "Orange", "Banana") //Adding element at the end using append function
	fmt.Print("Fruite list is: %T\n", fruitList)

	fruitList = append(fruitList[1:])

}
