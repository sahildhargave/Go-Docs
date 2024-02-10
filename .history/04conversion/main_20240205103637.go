//😁😂🤣😃😄
// conversion in go like string to int
// as in previous example 03 variable its giving the string if number is added

package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Welcome to our pizza app");
	fmt.Println("Please rate our pizza between 1 to 5");

	reader := bufio.NewReader(os.Stdin);

	imput, _ := reader.ReadString('\n')
   fmt.Println("Thanks for rating",)
}