package model

import (
	"Test/mySqldb"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = mySqldb.MySqlDb

type ReturnType struct {
	Msg  string
	Data interface{}
}
