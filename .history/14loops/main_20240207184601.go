package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in GOlang")

	days := []string{
		"Sun",
		"Mon",
		"Tue",
		"Wed",
		"Thu",
		"Fri",
		"Sat",
	}

	fmt.Println("Dayes in weeks are:", days)
	fmt.Printf("Dayes in weeks are: %+v", days)

	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i]) //accessing array using index
	//}
	//
	//	for i := range days {
	//		fmt.Println(days[i])
	//	}
	//}

	//for _, day := range days {
	//	fmt.Printf("index is and value is %v\n", day)
	//}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 5 {
			//break
			continue
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}
}
