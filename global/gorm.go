package global

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Gorm() {
	dns := "root:liufuling@tcp(14.103.50.35:3306)/table_a?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	if err != nil {
		log.Println(err)
	}

	DB = db
	db1, _ := DB.DB()
	db1.SetMaxIdleConns(10)
	db1.SetMaxOpenConns(100)
	db1.SetConnMaxIdleTime(time.Hour)

	err1 := DB.AutoMigrate()
	if err1 != nil {
		log.Println("失败")
	}
}
