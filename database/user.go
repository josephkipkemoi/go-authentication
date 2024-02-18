package database

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Passcode int    `json:"passcode"`
}

func (u User) SaveUser(db *sql.DB) (User, error) {
	insertUserSQL := `INSERT INTO users(email, password, passcode) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // prepare statement // good to avoid sql injection
	if err != nil {
		log.Fatal(err.Error())
	}

	// num := strconv.Itoa(u.PhoneNumber)
	// // Confirm number has set number of digits (9) before saving
	// if len(num) != 12 {
	// 	return User{}, errors.New("invalid phone number format. Expects 12 digits.")
	// }

	// Save user resource to DB
	res, err := statement.Exec(u.Email, u.Password, u.Passcode)
	if err != nil {
		return User{}, errors.New("user already registered. login to continue")
	}

	id, _ := res.LastInsertId()

	u.Id = id

	return u, nil
}

func (u User) GetUserID(email string) (int, string) {
	getUserSQL := `SELECT id FROM users WHERE phone_number = ?`
	row, err := DB.Query(getUserSQL, email)
	if err != nil {
		return 0, err.Error()
	}

	var id int
	for row.Next() {
		row.Scan(&id)
	}
	return id, ""
}

func (u User) GetUsers() []User {
	var usr User
	getUsersSQL := `SELECT * FROM users`
	row, err := DB.Query(getUsersSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	// var id int64
	// var email string
	// var password string
	// var passcode int

	var users []User
	for row.Next() {
		row.Scan(&usr.Email, &usr.Password, &usr.Passcode)
		users = append(users, usr)
	}
	return users

}

func (u User) AuthenticateUser(phone int, password string, db *sql.DB) (User, bool) {
	var email string

	selectQuery := `SELECT phone_number FROM users WHERE phone_number = ? AND password = ?`
	if err := db.QueryRow(selectQuery, phone, password).Scan(&email); err != nil {
		u.GetUsers()
		return User{}, false
	}

	return User{Email: email}, true
}
