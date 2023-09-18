package main

import (
	"fmt"
	"github.com/adarsh-jaiss/Go-Auth/pkg"
)

func main()  {
	// connecting to DB
	err := pkg.OpenDB()
	if err!= nil{
		fmt.Printf("database not found %v\n", err)
	}
	pkg.CreateTable()
	defer pkg.CloseDB()

}