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
	os.Remove("./database/files/pinacle.db") // delete file to avoid duplication
	log.Println("Creating pinacle database...")

	file, err := os.Create("./database/files/pinacle.db") // create SQLite file
	checkErr(err)
	file.Close()
	log.Println("pinacle.db created...")

	pinacleDb, _ := sql.Open("sqlite3", "./database/files/pinacle.db") // open the created pinacle db file
	// defer pinacleDb.Close()                                            // defer closing the database
	createUsersTable(pinacleDb)

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
