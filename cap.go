package main

import (
	"fmt"
)

func main()  {
	s := []int{1, 3, 5, 7, 9}
	fmt.Printf("cap: %d, len: %d\n", cap(s), len(s))

	ss := s[1:3]
	ss[0] = 111
	fmt.Printf("cap: %d, len: %d\n", cap(s), len(s))

	ss = append(ss, 10, 11, 12)
	fmt.Printf("cap: %d, len: %d\n", cap(ss), len(ss))
	fmt.Printf("cap: %d, len: %d\n", cap(s), len(s))
	s[1] = 333
	fmt.Println(s)
	fmt.Println(ss)
}