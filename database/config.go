package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnentDB() {
	pinacleDb := createDatabase()

	DB = pinacleDb
}

func createDatabase() *sql.DB {
	var fp string = "./database/files/pinacle.db"
	ok := checkFileExists(fp)

	if !ok {
		file, err := os.Create(fp) // create SQLite file
		checkErr(err)
		file.Close()
		log.Println("pinacledb.db created...")
	}

	pinacleDb, _ := sql.Open("sqlite3", fp) // open the created pinacle db file

	if !ok {
		createUsersTable(pinacleDb)
	}

	log.Println("pinacledb running...")

	return pinacleDb
}

func createUsersTable(db *sql.DB) {
	createUsersTable := `CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email STRING NOT NULL,
		password STRING NOT NULL,
		passcode STRING NOT NULL
	);` // SQL statement to create users table

	log.Println("Creating users table...")

	statement, err := db.Prepare(createUsersTable) // prepare SQL statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // execute SQL statement

	log.Println("users table created!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkFileExists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil
}
