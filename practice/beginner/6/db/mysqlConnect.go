package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	driver string = "mysql"
	user   string = "root"
	pass   string = "P@ssW0rd"
	host   string = "localhost"
	port   string = "3306"
	name   string = "beginner_6"
)

func InitDB(name string) (*sql.DB, error) {
	db, err := sql.Open(driver, user+":"+pass+"@tcp("+host+":"+port+")/"+name)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	return db, nil
}
