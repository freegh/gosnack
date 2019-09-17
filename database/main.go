package main

import (
	"fmt"
	"database/sql"
	"log"
	_"github.com/lib/pq" //_บอกว่าไม่ได้ใช้ที่นี่ แต่มีใช้
	"os"
)

func createTable(db *sql.DB, err error)  {
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

func dropTable(db *sql.DB)  {
	dropTb := `
	DROP TABLE todos;
	`
	_, err := db.Exec(dropTb)

	if err != nil {
		log.Fatal("can't drop table", err)
	}

	fmt.Println("drop table success")
}

func insertTable(db *sql.DB)  {
	dropTb := `
	INSERT INTO todos (id, title, status)
	VALUES (1, "KNOT", "GOD");
	`
	_, err := db.Exec(dropTb)

	if err != nil {
		log.Fatal("can't drop table", err)
	}

	fmt.Println("insert table success")
}

func updateTable(db *sql.DB)  {
	dropTb := `
	UPDATE todos
	SET title = "Nus", status = "Satan"
	WHERE id = 1;
	`
	_, err := db.Exec(dropTb)

	if err != nil {
		log.Fatal("can't drop table", err)
	}

	fmt.Println("update table success")
}

func deleteTable(db *sql.DB)  {
	dropTb := `
	DELETE FROM todos WHERE id = 1;
	`
	_, err := db.Exec(dropTb)

	if err != nil {
		log.Fatal("can't drop table", err)
	}

	fmt.Println("insert table success")
}

func main()  {
	url :=os.Getenv("DATABASE_URL")
	fmt.Println("url:", url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	_, table_check := db.Query("select * from todos;")
	if table_check == nil {
		fmt.Println("table exist")
	} else {
		createTable(db, err)
	}
	dropTable(db)

}