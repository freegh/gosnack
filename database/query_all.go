package main

import (
	"fmt"
	"database/sql"
	"log"
	_"github.com/lib/pq"
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct{
	ID int
	Title string
	Status string
}

var db *sql.DB
var err error
var todos []Todo

func connectDb()  {
	url :=os.Getenv("DATABASE_URL")
	fmt.Println("url:", url)
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
}

func queryAll(c *gin.Context) {
	t := Todo{}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos") 
	if err != nil {
		log.Fatal("can't prepare query all todos statement", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			log.Fatal("can't scan row into variable", err)
		}
		fmt.Println(&t.ID, &t.Title, &t.Status)
		todos = append(todos,t)
	}
	fmt.Println("query all todos success")
	fmt.Println(todos)
	c.JSON(http.StatusOK,todos)
}

func postAll(c *gin.Context)  {
	t := Todo{}
	c.JSON(200, gin.H{
		"message": "POST",
	})
	if err := c.ShouldBindJSON(&t); err != nil { //ของใส่ t
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, t)

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2) RETURNING id", t.Title, t.Status)
	err = row.Scan(&t.ID)
	if err != nil {
		fmt.Println("can't scan id", err)
		return
	}
	fmt.Println("insert todo success", todos)
}

func updateAll(c *gin.Context)  {
	id := c.Param("id")
	
	stmt, err := db.Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statement update", err)
	}
	if _, err := stmt.Exec(id, "inactive"); err != nil {
		log.Fatal("error execute update ", err)
	}
	fmt.Println("update todo success", todos)
}

func main()  {
	r := gin.Default()
	
	connectDb()
	r.POST("/ping",postAll)
	r.GET("/ping",queryAll)
	r.PUT("/ping/:id",updateAll)
	r.Run(":1234")
	
	defer db.Close()
}