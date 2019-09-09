package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

// func authMiddleware(handler func(*gin.Context)) func(*gin.Context)  {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		if token != "Bearer token123" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "unautorized."})
// 			return
// 		}
//		handler(c)
// 	}
// }

func authMiddleware(c *gin.Context)  {
	fmt.Println("This is a middlewear")
	token := c.GetHeader("Authorization")
	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized."})
		c.Abort()
		return
	}

	c.Next()

	fmt.Println("after in middleware")
	
}

func getAllTodo (c *gin.Context)  {
	token := c.GetHeader("Authorization")
	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unautorized."})
		return
	}

	t := Todo{}
	c.JSON(200, gin.H{
		"message": "POST",
	})
	if err := c.ShouldBindJSON(&t); err != nil { //ของใส่ t
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = len(todos)+1
	todos = append(todos, t)
}

func main() {
	r := gin.Default()
	r.Use(authMiddleware)
	r.GET("/ping", func (c *gin.Context)  {
		//c.JSON(200, todos)
		title := c.Query("title")
		fmt.Println(title)
		result := []Todo{}
		for _, v := range todos {
			if v.Title == title {
				result = append(result, v)
			}
		}
		c.JSON(200, result)
	})
	r.GET("/ping/:id", func (c *gin.Context)  {
		id := c.Param("id")
		result := []Todo{}
		for _, v := range todos {
			if strconv.Itoa(v.ID) == id {
				result = append(result, v)
			}
		}
		c.JSON(200, result)
	})
	r.POST("/ping", getAllTodo)
	r.Run(":1234")
}