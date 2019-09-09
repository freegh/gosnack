package main

import (
	"fmt"
	"database/sql"
	"log"
	_"github.com/lib/pq" //_บอกว่าไม่ได้ใช้ที่นี่ แต่มีใช้
	"os"
)

func main()  {
	url :=os.Getenv("DATABASE_URL")
	fmt.Println("url:", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`
	_, err =db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")
}