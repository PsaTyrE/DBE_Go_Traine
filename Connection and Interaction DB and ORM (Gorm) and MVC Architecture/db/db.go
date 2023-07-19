package db

import (
	"database/sql"
	"myApp/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var err error

func Init() {
	conf := config.GetConfig()

	conn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	// username:password@protocol(address)/dbname

	db, err = sql.Open("mysql", conn)
	if err != nil {
		panic("koneksi error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
