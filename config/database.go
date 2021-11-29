package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Tes() {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err=", err)
	}
	defer db.Close()

}
