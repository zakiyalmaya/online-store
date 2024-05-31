package model

import "time"

type CustomerEntity struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	Address     string    `db:"address"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
