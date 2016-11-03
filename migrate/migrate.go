package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  //? find out what is _ meaning
	"github.com/techusa786/nmhutil"
)

func main() {
	db, err := sql.Open("mysql", "root:wizard786@tcp(127.0.0.1:3306)/gotest")
	nmhutil.CheckError(err)

	defer db.Close()

	// let us check if connection is available
	err = db.Ping()
	nmhutil.CheckError(err)

	// drop table first, if exists
	query := "DROP TABLE IF EXISTS person"
	stmt, err := db.Prepare(query)
	nmhutil.CheckError(err)
	_, err = stmt.Exec()
	nmhutil.CheckError(err)

	// create table
	query = "CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));"
	stmt, err = db.Prepare(query)
	nmhutil.CheckError(err)
	_, err = stmt.Exec()
	nmhutil.CheckError(err)

	// if we reached here, query executed properly
	fmt.Println("person table is create successfully")
}