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

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 245
	highScore[2] = 254
	highScore[3] = 233

	highScore = append(highScore, 555, 334, 232)

	fmt.Println(highScore)
}
