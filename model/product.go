package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductEntity struct {
	ID            int             `db:"id"`
	Name          string          `db:"name"`
	Price         decimal.Decimal `db:"price"`
	StockQuantity int             `db:"stock_quantity"`
	CategoryID    int             `db:"category_id"`
	Description   string          `db:"description"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}

type ProductRequest struct {
	Name          string  `json:"name" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	StockQuantity int     `json:"stock_quantity" validate:"required"`
	CategoryID    int     `json:"category_id" validate:"required"`
	Description   string  `json:"description,omitempty"`
}

type ProductResponse struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
	Category      string  `json:"category"`
	Description   string  `json:"description"`
}
