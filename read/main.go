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

	ambil := []config.Student{}

	db.Where("age < ?", "19").Find(&ambil)

	for _ , data := range ambil {
		fmt.Println("name " , data.Name)
	}

	search := "%" + "Jo" + "%"
	db.Where("name LIKE ?", search).Find(&ambil)

	for _, dataku := range ambil {
		fmt.Println("jeneng", dataku.Name)
	}
}