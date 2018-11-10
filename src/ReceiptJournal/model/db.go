package model

import (
	"database/sql"
)

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}

var Hostadd string = "hostadd"
var Dbname string = "dbname"
var Dbuser string = "dbusername"
var Dbpass string = "dbpasswor"
