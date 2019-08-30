package main

import (
	"fmt"
)

type student struct {
	name string
	age int
}

var students = map[int]student{
	1: {name: "Earth", age: 12},
	2: student{
		name: "Chin", age: 11,
	},
}

func main()  {
	if s, ok := students[3]; ok {
		fmt.Println(ok)
	} else {
		fmt.Println("som $ #v\n", s)
	}
	fmt.Println(students)
}