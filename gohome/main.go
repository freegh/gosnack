package main

import (
	"fmt"
	"gohome/user" //in GOPATH/src
	myC "gohome/a/b/c"
	"gohome/c"

	"github.com/AnuchitO/say"
)

func Something(new, word string) {
	fmt.Println("say something is ", word)
}

func main() {
	u := user.User{Name: "Jojo", Email: "Jo@jmail.com"}
	c.Show(u)
	myC.Show(u)

	say.Something("KNOT")
}