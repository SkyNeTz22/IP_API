package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var config Config

func InitialMigration() {
	file, err := os.ReadFile("config.json")

	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &config)

	cfg := mysql.Config{
		User:   config.MySql.Username,
		Passwd: config.MySql.Password,
		Net:    config.MySql.Network,
		Addr:   config.MySql.HostName,
		DBName: config.MySql.DbName,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}

}
