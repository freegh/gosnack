package main

import (
	"fmt"
)

func WordCount(s string) map[string]int{
	result := make(map[string]int)
	var word string
	for i := 0; i <= len(s); i++ {
		
		if (i == len(s) || s[i:i+1] == " "){
			result[word] += 1
			word = ""
		}else{
			word += s[i:i+1]
		}
		
	}
	 
	return result
}

func main()  {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}