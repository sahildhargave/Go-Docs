//😁😂🤣😃😄
// conversion in go like string to int
// as in previous example 03 variable its giving the string if number is added

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of the rating %T", input)

	numRating, err := strconv.ParseFloat(input, 64);

	if err != nil {
     fmt.Println(err)
	}else{
		fmt.Println("Original Rating of the user", numRating);
		
	}

}
