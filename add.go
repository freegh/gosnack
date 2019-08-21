package main

import "fmt"

func main(){
	fmt.Println(add(1-3,2))
	fmt.Print(swap(1,2))
}

func add(x, y int) int {
	return x+y
}

func swap(x ,y int) (int,int) {
	return y,x
}