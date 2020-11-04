package main

import (
	"github.com/pengdafu/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:pdf0824@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{Logger: &logger.Logger{}})
	if err != nil {
		panic("连接错误")
	}
}
