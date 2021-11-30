package model

import (
	"Test/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = config.MysqlDb
