package models

import (
	"errors"
	"time"

	"example.com/online-store/db"
	"example.com/online-store/utils"
	"github.com/guregu/null/v5"
)

type User struct {
	ID          int64       `json:"id"`
	Avatar      null.String `json:"avatar"`
	FirstName   null.String `json:"first_name"`
	LastName    null.String `json:"last_name"`
	Username    null.String `json:"username"`
	Email       null.String `json:"email" validate:"required,email"`
	Password    null.String `json:"password" validate:"required"`
	BirthOfDate NullTime    `json:"birth_of_date"`
	PhoneNumber null.String `json:"phone_number"`
	CreatedAt   NullTime    `json:"created_at"`
	DeletedAt   NullTime    `json:"deleted_at"`
}

func GetAllUsers() ([]User, error) {
	var users []User
	var query string = "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Avatar, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.BirthOfDate, &user.PhoneNumber, &user.CreatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(userID int64) (User, error) {
	var query = "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.ID, &user.Avatar, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.BirthOfDate, &user.PhoneNumber, &user.CreatedAt, &user.DeletedAt)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) Save() error {
	var query string = `
		INSERT INTO users(avatar, first_name, last_name, username, email, password, birth_of_date, phone_number, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	u.Password.String, err = utils.HashPassword(u.Password.String)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Avatar, u.FirstName, u.LastName, u.Username, u.Email, u.Password, u.BirthOfDate, u.PhoneNumber, time.Now())

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	return err

}

func (u *User) Update() error {
	var query = `
	UPDATE users
	SET avatar = ?, first_name = ?, last_name = ?, username = ?, email =?, password = ?, birth_of_date = ?, phone_number = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Avatar, u.FirstName, u.LastName, u.Username, u.Email, u.Password, u.BirthOfDate, u.PhoneNumber, u.ID)

	return err
}

func (u *User) Delete() error {
	var query = `
	UPDATE users
	SET deleted_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), u.ID)

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CompareHashPassword(u.Password.String, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
