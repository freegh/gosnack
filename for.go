package main

import "fmt"

func main () {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	A := 0
	for A < 10 {
		i := 1
		A += i
	}
	names := []string{"game","bank","samui","dog","jiew"}
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println(sum)
	fmt.Println(A)
}