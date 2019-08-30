package main

import (
	"fmt"
)

func main() {
	i := 42
	fmt.Println("i:", i)

	var p *int
	fmt.Println("p:", p)
	fmt.Println("i: %p\n", &i)

	p = &i
	fmt.Println("p:", p)

	i = 55
	fmt.Println("p:", *p)
	fmt.Println("p:", p)

	*p = 33
	fmt.Println("new i :", i)
}