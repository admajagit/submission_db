package config

import (
	"fmt"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id        uint
	Name      string
	Age       uint8
	CreatedAt time.Time
	UpdateAT  time.Time
	DeletedAt gorm.DeletedAt
}

func Newdatabase() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error cuy")
	}

	db.AutoMigrate(&Student{})

	return db

}
