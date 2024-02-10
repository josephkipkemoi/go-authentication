package database

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	PhoneNumber     int    `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u User) SaveUser(db *sql.DB) (User, error) {
	insertUserSQL := `INSERT INTO users(phone_number, password) VALUES (?, ?)`
	statement, err := db.Prepare(insertUserSQL) // prepare statement // good to avoid sql injection
	if err != nil {
		log.Fatal(err.Error())
	}
	// Add country prefix to phonenumber, receive 700545727 save 254700545727
	fmt.Println(u.PhoneNumber)
	_, err = statement.Exec(u.PhoneNumber, u.Password)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	return u, nil
}

func (u User) GetUsers(db *sql.DB) {
	getUsersSQL := `SELECT * FROM users`
	row, err := db.Query(getUsersSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	for row.Next() {
		var id int
		var phone_number int
		var password string
		row.Scan(&id, &phone_number, &password)
		log.Println(id, phone_number, password)
	}
}

func (u User) AuthenticateUser(phone int, password string, db *sql.DB) (User, bool) {
	var num int
	if err := db.QueryRow("SELECT phone_number FROM users WHERE phone_number = ? AND password = ?", phone, password).Scan(&num); err != nil {
		return User{}, false
	}
	return User{PhoneNumber: num}, true
}
