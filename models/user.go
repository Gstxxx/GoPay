package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CPF        string  `json:"cpf"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Balance    float64 `json:"balance"`
	IsMerchant bool    `json:"is_merchant"`
	Token      string  `json:"token"`
}

var Users []User
var NextUserID int = 1

var db *sql.DB

func InitDatabase() {
	dataSourceName := fmt.Sprintf("root:sua_senha_secreta@tcp(localhost:3306)/gopay_mysql")
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
