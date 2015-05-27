package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DBH *sql.DB
)

func createDBFile() {
	//path := path.Dir("." + string(filepath.Separator) + DB_DIR)
	path := "." + string(filepath.Separator) + DB_DIR
	fmt.Printf("path(%T): %v", path, path)

	// create database directory
	if os.MkdirAll(path, 0755) != nil {
		log.Fatal("couldn't create database directory ./" + DB_DIR + "/")
	} else {
		fmt.Println("created database directory " + path)
	}

	// create database file
	_, err := os.Create(DB_PATH)
	if err != nil {
		log.Fatal("couldnt create database file " + DB_PATH)
	}

}

func initDB() {
	// check if database exists
	fmt.Println("opening database " + DB_PATH + "?cache=shared&mode=wrc... ")
	if _, err := os.Stat(DB_PATH); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("database file doesn't exist, creating it...")
			createDBFile()
		}
	}

	var err error
	DBH, err = sql.Open("sqlite3", "file:"+DB_PATH+"?cache=shared&mode=rwc")

	// TODO: handle this properly
	if err != nil {
		fmt.Println("exit 1")
		log.Fatal(err)
	}
	if DBH == nil {
		fmt.Println("exit 2")
		log.Fatal(err)
	}

	fmt.Println("done!")
}

func createTables() {

	// TODO: move this!

	fmt.Printf("creating tables and triggers if they don't exist... ")
	SQLiteStatements := []string{
		`CREATE TABLE if not exists markets (
                        uid integer not null primary key AUTOINCREMENT, 
                        unixtime int not null, 
                        cryptsy TEXT,
			poloniex TEXT,
			bittrex TEXT,
			daily_volume TEXT,
			USD TEXT	
                )`,
	}

	for _, v := range SQLiteStatements {
		tx, err := DBH.Begin()
		if err != nil {
			fmt.Println("exit 16")
			log.Fatal(err)
		}

		stmt, err := tx.Prepare(v)
		if err != nil {
			fmt.Println("exit 16")
			log.Fatal(err)
		}

		_, stmt_err := stmt.Exec()
		if stmt_err != nil {
			fmt.Printf("%q: %s\n", err, v)
			fmt.Println("exit 17")
			log.Fatal(stmt_err)
		}
		tx.Commit()
		stmt.Close()
	}
	fmt.Println(" done!")

}
