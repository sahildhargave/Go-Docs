package main

import "fmt"

func main() {
	fmt.Println("Welcome to info of slice")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitList)

	fruitList = append(fruitList, "Orange", "Banana") //Adding element at the end using append function
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList) //Removing first element from list and adding it after last element

	highScore := make
}

