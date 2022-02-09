package app

import (
	"database/sql"
	"fmt"
	"giricorp/belajar-go-restfull-api/helper"
	"time"
)

type DBMain struct {
	Username  string
	Password  string
	Domain    string
	Port      int
	DBName    string
	ParseTime bool
}

type DBClient struct {
	Username  string
	Password  string
	Domain    string
	Port      int
	DBName    string
	ParseTime bool
}

func BuildStringParamsMain(dbc *DBMain) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", dbc.Username, dbc.Password, dbc.Domain, dbc.Port, dbc.DBName, dbc.ParseTime)
}

func BuildStringParamsClient(dbc *DBClient) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t", dbc.Username, dbc.Password, dbc.Domain, dbc.Port, dbc.DBName, dbc.ParseTime)
}

func NewDBMain() *sql.DB {
	dbMain := DBMain{
		Username:  "root",
		Password:  "root",
		Domain:    "localhost",
		Port:      3307,
		DBName:    "golang_rest",
		ParseTime: true,
	}
	db, err := sql.Open("mysql", BuildStringParamsMain(&dbMain))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}

func (dbClient *DBClient) NewDBClient() *sql.DB {
	db, err := sql.Open("mysql", BuildStringParamsClient(dbClient))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
