package main

import (
 "errors"
 "fmt"
)

type BusinessErr struct {
 Code     string
 Message  string
 Severity string
}

func (b BusinessErr) Error() string {
 return fmt.Sprintf("business error: code=%s, msg:%s, serverity:%s", b.Code, b.Message, b.Severity)
}

func PrintError(err error) {
 // e := err.Error()
 fmt.Println(err)
}

func DoSomething() error {
 // ....
 return BusinessErr{
  Code: "E002",
  //....
 }
}

func main() {
 err := errors.New("my error")

 bizErr := BusinessErr{
  Code:     "E001",
  Message:  "account not found",
  Severity: "3",
 }

 PrintError(err)
 PrintError(bizErr)
}