package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	user   string = "root"
	pass   string = "P@ssW0rd"
	host   string = "localhost"
	port   string = "3306"
	name   string = "beginner_6"
)

func InitDB(name string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(user+":"+pass+"@tcp("+host+":"+port+")/"+name), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	return db, nil
}
