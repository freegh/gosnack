package main

import "fmt"

const Pi = 3.14
const (
 Zero = iota-10+7*2/3
 One
 Two
 Three
 Four
 Five
)

func main(){
	const world = "gg"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)

	fmt.Println(Zero)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
	fmt.Println(Five)
}