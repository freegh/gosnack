package main

import "fmt"

func fibonacci() (func() int){
	A := 0
	B := 1
	return func() int {
		A, B = B, A+B
		return A
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}