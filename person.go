package main

import (
	"fmt"
)

type person struct{
	name string
	shoes []string
}

func (pers person) Walk() {
	fmt.Println("walking ", pers.name)
}
func (pers person) Eat() {
	fmt.Println("eating ", pers.name)
}
func (pers person) Geeting() {
	fmt.Println("hello ",pers.name)
}
func (pers person) Getter() string{
	return pers.name
}
func (pers *person) Setter(newName string) {
	pers.name = newName
	pers.shoes[0] = "Z"

}

func main()  {
	person1 := person{name:"Knot", shoes:[]string{"A", "B"},}
	person1.Walk()
	fmt.Println(person1.Getter())
	person1.Eat()
	person1.Setter("Go")
	person1.Geeting()
	fmt.Println(person1)
}