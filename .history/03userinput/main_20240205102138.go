//😁😂🤣😃😄

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	//😁😂🤣😃 comma ok // err ok
	// not having try and catch
	input, _ := reader.ReadString('\n')
	rating,_ := reader.Ne
	fmt.Println("Thanks for giving name,", input)
	fmt.Printf("Type of this name is %T,", input)
}
