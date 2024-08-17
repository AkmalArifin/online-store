package models

import (
	"time"

	"example.com/online-store/db"
	"github.com/guregu/null/v5"
)

type Address struct {
	ID           int64       `json:"id"`
	UserID       int64       `json:"user_id" validate:"required"`
	Title        null.String `json:"title"`
	AddressLine1 null.String `json:"address_line_1"`
	AddressLine2 null.String `json:"address_line_2"`
	Country      null.String `json:"country"`
	City         null.String `json:"city"`
	PostalCode   null.String `json:"postal_code"`
	Landmark     null.String `json:"landmark"`
	PhoneNumber  null.String `json:"phone_number"`
	CreatedAt    NullTime    `json:"created_at"`
	DeletedAt    NullTime    `json:"deleted_at"`
}

func GetAllAddresses() ([]Address, error) {
	var addresses []Address
	var query string = "SELECT * FROM addresses"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address Address
		err = rows.Scan(&address.ID, &address.UserID, &address.Title, &address.AddressLine1, &address.AddressLine2, &address.Country, &address.City, &address.PostalCode, &address.Landmark, &address.PhoneNumber, &address.CreatedAt, &address.DeletedAt)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func GetAddressById(addressID int64) (Address, error) {
	var query = "SELECT * FROM addresses WHERE id = ?"
	row := db.DB.QueryRow(query, addressID)

	var address Address
	err := row.Scan(&address.ID, &address.UserID, &address.Title, &address.AddressLine1, &address.AddressLine2, &address.Country, &address.City, &address.PostalCode, &address.Landmark, &address.PhoneNumber, &address.CreatedAt, &address.DeletedAt)

	if err != nil {
		return Address{}, err
	}

	return address, nil
}

func GetAddressByUserId(userID int64) ([]Address, error) {
	var addresses []Address
	var query string = "SELECT * FROM addresses WHERE user_id = ?"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address Address
		err = rows.Scan(&address.ID, &address.UserID, &address.Title, &address.AddressLine1, &address.AddressLine2, &address.Country, &address.City, &address.PostalCode, &address.Landmark, &address.PhoneNumber, &address.CreatedAt, &address.DeletedAt)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (a *Address) Save() error {
	var query string = `
	INSERT INTO addresses(user_id, title, address_line_1, address_line_2, country, city, postal_code, landmark, phone_number, created_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	a.CreatedAt.SetValue(time.Now())
	result, err := stmt.Exec(a.UserID, a.Title, a.AddressLine1, a.AddressLine2, a.Country, a.City, a.PostalCode, a.Landmark, a.PhoneNumber, a.CreatedAt)

	if err != nil {
		return nil
	}

	a.ID, err = result.LastInsertId()
	return err
}

func (a *Address) Update() error {
	var query = `
	UPDATE addresses
	SET title = ?, address_line_1 = ?, address_line_2 = ?, country = ?, city = ?, postal_code = ?, landmark = ?, phone_number = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Title, a.AddressLine1, a.AddressLine2, a.Country, a.City, a.PostalCode, a.Landmark, a.PhoneNumber, a.ID)

	return err
}

func (a *Address) Delete() error {
	var query = `
	UPDATE addresses
	SET deleted_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), a.ID)

	return err
}
