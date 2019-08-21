package main

import "fmt"

func main(){
	defer fmt.Println("my defer 1")
	defer fmt.Println("my defer 2")
	defer fmt.Println("my defer 3")

	for i := 0; i < 10; i++ {
		defer func(){
			fmt.Println(i)
		}()
	}

	fmt.Println("in main")
}