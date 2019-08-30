package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-q.Y)
}

// func (p Point) Distance(q Point) float64 {
// 	return math.Hypot(q.X-q.X, q.Y-q.Y)
// }

func (p Point) Distance() {
	fmt.Println("X", p.X, "Y", p.Y)
}

type ColorPoint struct {
	Point
	Color string
}

type Path []Point


func main()  {
	// p := Point{ X:1.3, Y:2.1}
	// q := Point{ 1, 2 }
	// fmt.Println(Distance(p,q))
	// fmt.Println(p.Distance(q))

	p := Point{2, 4}
	c := ColorPoint{Point: p, Color: "blue"}
	fmt.Println(c.X)
}