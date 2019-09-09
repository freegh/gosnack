package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Todo struct{
	ID string `json:"id"`
	Title string `json:"tittle"`
	Status string `json:"status"`
}

var todos []Todo

func todosHandler (w http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "GET" {
		//w.Write([]byte(`{"name": "knot"}`))

		b, err := json.Marshal(todos)
		w.Write(b)
		fmt.Printf("%T => %v \n %s \n", b, b, b)
		fmt.Println(err)

		w.Header().Set("Content-Type", "application/json")
		return
	}
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Todo{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Fprintf(w, "error: ", err)
		}
		fmt.Printf("body : % #v\n", t)

		fmt.Printf("body : %s\n", body)
		fmt.Fprintf(w, "OK")

		todos = append(todos, t)
		return
	}
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	err := http.ListenAndServe(":1234", nil)
	fmt.Println("after start")
	fmt.Println(err)
}