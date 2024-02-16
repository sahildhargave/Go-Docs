package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("world")
	greateNight("earth")

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func greateNight(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("Sahil Dhargave")
	}
}
