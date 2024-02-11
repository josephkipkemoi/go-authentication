package database

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type User struct {
	Id              int64  `json:"id"`
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

	num := strconv.Itoa(u.PhoneNumber)
	// Confirm number has set number of digits (9) before saving
	if len(num) != 12 {
		return User{}, errors.New("invalid phone number format. Expects 12 digits.")
	}
	// Add country prefix to phonenumber, receive 700545727 save 254700545727
	n, err := strconv.Atoi("254" + num)

	if err != nil {
		return User{}, err
	}

	// Save user resource to DB
	res, err := statement.Exec(n, u.Password)
	if err != nil {
		return User{}, errors.New("user already registered. login to continue.")
	}

	id, _ := res.LastInsertId()

	u.Id = id
	u.PhoneNumber = n

	return u, nil
}

func (u User) GetUser(db *sql.DB, phone_number int) (int, string) {
	getUserSQL := `SELECT id FROM users WHERE phone_number = ?`
	row, err := db.Query(getUserSQL, phone_number)
	if err != nil {
		return 0, err.Error()
	}

	var id int
	for row.Next() {
		row.Scan(&id)
	}
	return id, ""
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
	selectQuery := `SELECT phone_number FROM users WHERE phone_number = ? AND password = ?`
	if err := db.QueryRow(selectQuery, phone, password).Scan(&num); err != nil {
		return User{}, false
	}
	return User{PhoneNumber: num}, true
}
