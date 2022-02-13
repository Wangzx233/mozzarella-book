package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

var (
	//username = "root"
	//password = ""
	//host     = ""
	//port     = "3306"
	//dbname   = "mozzarella-book"
	username = "root"
	password = "root"
	host     = "localhost"
	port     = "3306"
	dbname   = "mozzarella_book"
)

func Init() {
	dsn := username + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	err = DB.AutoMigrate(&Book{}, &Images{}, &Cart{})
	if err != nil {
		log.Println(err)
		return
	}

}
