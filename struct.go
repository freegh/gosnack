package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{Y:1, X:2}
	fmt.Println(v)
	v = Vertex{1,2}
	fmt.Printf("v % #v\n", v)
	v.X = 5
	fmt.Println(v)
	v = Vertex{X:2}
	fmt.Println(v)

	p := &v
	p.Y=16
	fmt.Println(p)
	fmt.Println(*p)

	var pp *Vertex
	pp = &v
	pp.Y = 11
	fmt.Println(pp)
	fmt.Println(*pp)
}