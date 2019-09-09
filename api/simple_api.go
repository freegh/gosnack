package main

import (
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name": "knot"}`) //back quote
	w.Write(resp)
}

func main ()  {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":1234", nil)
	fmt.Println("after start")
	fmt.Println(err)
	log.Fatal(err) //kill server and gen log
}