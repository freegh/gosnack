package main

import (
	"fmt"
)

func main()  {
	var m map[string]string
	m = make(map[string]string)
	m["Knot"] = "Chayapol"
	fmt.Println(m["Knot"])

	
}