package model

import (
	"time"

	"github.com/shopspring/decimal"
	cartEnum "github.com/zakiyalmaya/online-store/constant/cart"
)

type CartEntity struct {
	ID         int             `db:"id"`
	CustomerID int             `db:"customer_id"`
	Status     cartEnum.Status `db:"status"`
	CreatedAt  time.Time       `db:"created_at"`
	UpdatedAt  time.Time       `db:"updated_at"`
	Items      []*CartItemEntity
}

type CartItemEntity struct {
	ID          int             `db:"id"`
	ProductID   int             `db:"product_id"`
	Quantity    int             `db:"quantity"`
	CartID      int             `db:"shopping_cart_id"`
	CreatedAt   time.Time       `db:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at"`
	Price       decimal.Decimal `db:"price"`
	ProductName string          `db:"product_name"`
}

type GetCartRequest struct {
	CustomerID int  `json:"customer_id" validate:"required"`
	Status     *int `json:"status,omitempty"`
}

type CreateCartRequest struct {
	CustomerID int                      `json:"customer_id" validate:"required"`
	Items      []*CreateCartItemRequest `json:"items" validate:"required"`
}

type CreateCartItemRequest struct {
	ProductID int             `json:"product_id" validate:"required"`
	Quantity  int             `json:"quantity" validate:"required"`
	Price     decimal.Decimal `json:"price"`
}

type CartResponse struct {
	ID         int                 `json:"id"`
	CustomerID int                 `json:"customer_id"`
	Status     string              `json:"status"`
	Items      []*CartItemResponse `json:"items"`
}

type CartItemResponse struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (c *CreateCartRequest) ToEntity() *CartEntity {
	cartItemEntity := make([]*CartItemEntity, len(c.Items))
	for i, item := range c.Items {
		cartItemEntity[i] = &CartItemEntity{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	return &CartEntity{
		CustomerID: c.CustomerID,
		Status:     cartEnum.CartStatusActive,
		Items:      cartItemEntity,
	}
}

func (c *CartEntity) ToResponse() *CartResponse {
	cartItemResponse := make([]*CartItemResponse, len(c.Items))
	for i, item := range c.Items {
		cartItemResponse[i] = &CartItemResponse{
			ID:          item.ID,
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price.InexactFloat64(),
		}
	}

	return &CartResponse{
		ID:         c.ID,
		CustomerID: c.CustomerID,
		Status:     c.Status.Enum(),
		Items:      cartItemResponse,
	}
}
