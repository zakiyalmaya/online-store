package model

import "time"

type CategoryEntity struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
