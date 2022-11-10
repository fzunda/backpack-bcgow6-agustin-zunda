package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	configDB := mysql.Config{
		User:   os.Getenv("USERNAME"),
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DATABASE"),
	}

	db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("No conecto")
		panic(err)
	}

	return db
}
