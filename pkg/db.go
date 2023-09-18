package pkg

import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func OpenDB() error  {
	var err error
	DB,err := sql.Open("postgres", "user=postgres dbname=auth password='' sslmode=disable")
	if err != nil {
		return err
	}
	
	fmt.Printf("Connection established %v",DB)
	return nil

	
}

func CloseDB() error {
	return DB.Close()
	 
}

func CreateTable()  {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS USERS(id SERIAL PRIMARY KEY NOT NULL, username TEXT NOT NULL,email TEXT NOT NULL,password TEXT NOT NULL )")
	if err != nil {
		log.Fatal(err)
	}
}