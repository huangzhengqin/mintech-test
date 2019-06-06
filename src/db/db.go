package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)



func OpenDB () *gorm.DB {
	db, err := gorm.Open("mysql", "root:luyun123@/test?charset=utf8")
	if err != nil {
		log.Fatalf("open the DB fail, err:%s", err)
	}

	fmt.Println("open db success. ")

	return db
}