package main

import (
	"fmt"
	"go-sql/config"
	
)

func main() {
	fmt.Println("conexction database")

	db := config.Newdatabase()

	if db == nil {
		fmt.Println("failed conection")
		return
	}

	fmt.Println("main function has been triggered")


	student1 := []config.Student{
		{
			Name: "John",
			Age: 25,
		},
		{
			Name: "Thomas",
			Age: 30,
		},
		{
			Name: "Jane",
			Age: 18,
		},
		{
			Name: "Joe",
			Age: 17,
		},
	
	}

	db.Create(&student1)
}