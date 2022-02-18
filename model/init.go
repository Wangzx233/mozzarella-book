package model

import (
	"github.com/go-redis/redis"
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
	password = "123456"
	host     = "localhost"
	port     = "3306"
	dbname   = "mozzarella_book"
)

func InitMysql() {
	dsn := username + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	err = DB.AutoMigrate(&Book{}, &Images{}, &Cart{}, &Value{}, &UserClickRecord{})
	if err != nil {
		log.Println(err)
		return
	}

}

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		log.Println(err)
	}
}
