package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
 CREATE TABLE IF NOT EXISTS "newsfeed" (
	 "ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	 "Content" TEXT
 )
*/
func main() {
	db, err := sql.Open("sqlite3", "./newsfeed.db")
	checkErr(err)
	// stmt, _ := db.Prepare(`
	// INSERT INTO newsfeed (content) values (?)
	// `)
	// stmt.Exec("Yello Go Lovers")
	stmt, errdb := db.Prepare(`
	 CREATE TABLE "newsfeed" (
		"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Content" TEXT
	`)
	stmt.Exec()
	checkErr(errdb)
	// fmt.Println(errdb)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
