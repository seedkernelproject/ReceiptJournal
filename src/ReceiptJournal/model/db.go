package model

import (
	"database/sql"
)

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}

var Hostadd string = "alien-clover-220213:asia-southeast1:receiptjournal"
var Dbname string = "postgres"
var Dbuser string = "postgres"
var Dbpass string = "0811635143"
