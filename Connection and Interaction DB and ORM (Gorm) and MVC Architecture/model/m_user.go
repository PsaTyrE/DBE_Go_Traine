package model

import (
	"database/sql"
	"fmt"
	"myApp/db"
	"myApp/helper"
)

type User struct {
	iD       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	conn := db.CreateCon()

	query := "SELECT * FROM user WHERE username=?"

	err := conn.QueryRow(query, username).Scan(
		&obj.iD, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("username Not Found")
		return false, nil
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckpasswordHash(password, pwd)

	if !match {
		fmt.Println("hash password dosnt match")
		return false, err
	}

	return true, nil
}
