package main

import (
 "fmt"
 "os"
)

type errorWrite struct {
 file *os.File
 err  error
}

func (ew *errorWrite) Write(b []byte) error {
 if ew.err != nil {
  return ew.err
 }

 _, ew.err = ew.file.Write(b)
 return ew.err
}

func main() {
 f, err := os.Open("./name.txt")
 if err != nil {
  //...
 }
 ew := &errorWrite{file: f}
 ew.Write([]byte("a"))
 ew.Write([]byte("b"))
 ew.Write([]byte("c"))

 fmt.Println(ew.err)
}