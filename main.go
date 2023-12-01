package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tooltoys/db-tools/tools"
)

var mysqlConn *sqlx.DB

func init() {
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn")
	// conn, err := sqlx.Connect("postgres", "user=admin password=secret host=192.168.49.2 port=30303 dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConn = conn
}

func main() {
	m := tools.NewMulticolumnIndex(mysqlConn)
	m.AnalysisOrder("payment", "staff_id", "customer_id")
}
