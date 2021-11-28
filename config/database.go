package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db,err := gorm.Open("mysql",root:@tcp())
}
