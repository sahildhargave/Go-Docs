package main

import (
	"fmt"
	"os"
	"io"
)
//
//func main(){
//	fmt.Println(os.Args)
//	os.Exit(1)
//}

//package main

func main(){
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error",err)
		os.Exit(1)
	}
	io.Copy(os.Stdout,f)
}