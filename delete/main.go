package main

import (
	"fmt"
	"go-sql/config"
	
)

func main()  {
	fmt.Println("conexction database")

	db := config.Newdatabase()

	if db == nil {
		fmt.Println("failed conection")
		return
	}

	fmt.Println("main function has been triggered")
	




}