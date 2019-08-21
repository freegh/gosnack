package main

import "fmt"

func main(){
	var i int = 24
	var	f float32 = float32(i)
	fmt.Printf("i: %#v, f: %T", i ,f)
}