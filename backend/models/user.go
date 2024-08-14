package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"example.com/online-store/db"
	"github.com/go-sql-driver/mysql"
	"github.com/guregu/null/v5"
)

type User struct {
	ID          int64       `json:"id" validate:"required"`
	Avatar      null.String `json:"avatar"`
	FirstName   null.String `json:"first_name"`
	LastName    null.String `json:"last_name"`
	Username    null.String `json:"username" validate:"required"`
	Email       null.String `json:"email" validate:"required,email"`
	Password    null.String `json:"password" validate:"required"`
	BirthOfDate NullTime    `json:"birth_of_date" time_format:"2006-01-02"`
	PhoneNumber null.String `json:"phone_number"`
	CreatedAt   NullTime    `json:"created_at"`
	DeletedAt   NullTime    `json:"deleted_at"`
}

type NullTime struct {
	mysql.NullTime
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}

func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == 'n' {
		nt.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &nt.Time); err != nil {
		return fmt.Errorf("null: couldn't unmarshal JSON: %w", err)
	}

	nt.Valid = true
	return nil
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

	result, err := stmt.Exec(u.Avatar, u.FirstName, u.LastName, u.Username, u.Email, u.Password, u.BirthOfDate, u.PhoneNumber, time.Now())

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	return err

}

func GetAllUsers() ([]User, error) {
	var users []User
	var query string = "SELECT id, avatar, first_name, last_name, username, email, password, birth_of_date, phone_number, created_at, deleted_at FROM users"
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
