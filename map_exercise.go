package main

import (
	"fmt"
)

type StudentList struct{
	name string
	class string
	age int
}

func main()  {
	var student map[int]StudentList
	student = make(map[int]StudentList)

	student[1] = StudentList{"aa","A",23}
	fmt.Println(student[1])
}