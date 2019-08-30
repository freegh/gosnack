package main

import (
	"fmt"
)

type rectangle struct {
	width int
	length int
}

func area(rec rectangle) int {
	return rec.width * rec.length
}

func (rec rectangle) area() int {
	fmt.Println("method area")
	return rec.width * rec.length
}

func main()  {
	rec := rectangle{3, 4}
	a := area(rec)
	fmt.Println(a)

	b := rec.area()
	fmt.Println(b)
}