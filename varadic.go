package main

import "fmt"

func say(s string, ss ...string) {
 fmt.Println("say", s, ss)
 fmt.Println("OK")
}

func main() {
 s := []string{"a", "b", "c", "d", "e"}
 say("a", "b", "c", "d")
 say("a")
 say("test", s...)
 say(s[0], s[1:]...)
}