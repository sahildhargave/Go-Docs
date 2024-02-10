//😁😂🤣😃😄
// conversion in go like string to int
// as in previous example 03 variable its giving the string if number is added

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of the rating %T", input)
	// as this line give error becaue of the input hence

	//	numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Original Rating of the user", numRating)
		fmt.Println("Added 1 to your rating:", numRating+1)
		fmt.	printf("Type of the rating %T", numRating)Rating)
	}

}
