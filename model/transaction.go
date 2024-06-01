package model

import (
	"time"

	"github.com/shopspring/decimal"
	transactionEnum "github.com/zakiyalmaya/online-store/constant/transaction"
	"github.com/zakiyalmaya/online-store/utils"
)

type TransactionEntity struct {
	ID             int                    `db:"id"`
	IdempotencyKey string                 `db:"idempotency_key"`
	CustomerID     int                    `db:"customer_id"`
	CartID         int                    `db:"shopping_cart_id"`
	Status         transactionEnum.Status `db:"status"`
	TotalAmount    decimal.Decimal        `db:"total_amount"`
	PaymentMethod  transactionEnum.Method `db:"payment_method"`
	CreatedAt      time.Time              `db:"created_at"`
	UpdatedAt      time.Time              `db:"updated_at"`
	Details        []*TransactionDetailEntity
}

type TransactionDetailEntity struct {
	ID            int             `db:"id"`
	TransactionID int             `db:"transaction_id"`
	ProductID     int             `db:"product_id"`
	ProductName   string          `db:"product_name"`
	Quantity      int             `db:"quantity"`
	Price         decimal.Decimal `db:"price"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}

type TransactionRequest struct {
	CustomerID    int                    `json:"customer_id" validate:"required"`
	CartID        int                    `json:"shopping_cart_id" validate:"required"`
	PaymentMethod transactionEnum.Method `json:"payment_method" validate:"required"`
}

type TransactionResponse struct {
	ID             int                          `json:"id"`
	IdempotencyKey string                       `json:"idempotency_key"`
	CustomerID     int                          `json:"customer_id"`
	CartID         int                          `json:"shopping_cart_id"`
	Status         string                       `json:"status"`
	TotalAmount    float64                      `json:"total_amount"`
	PaymentMethod  string                       `json:"payment_method"`
	Details        []*TransactionDetailResponse `json:"transaction_details"`
}

type TransactionDetailResponse struct {
	ID          int     `json:"id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (c *CartEntity) ToTransactionEntity() *TransactionEntity {
	if len(c.Items) == 0 {
		return nil
	}

	var totalAmout float64
	details := make([]*TransactionDetailEntity, len(c.Items))
	for i, item := range c.Items {
		details[i] = &TransactionDetailEntity{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}

		totalAmout += item.Price.Mul(decimal.NewFromInt(int64(item.Quantity))).InexactFloat64()
	}

	return &TransactionEntity{
		IdempotencyKey: utils.GenerateUUID(),
		CustomerID:     c.CustomerID,
		CartID:         c.ID,
		Status:         transactionEnum.TransactionStatusInprogress,
		TotalAmount:    decimal.NewFromFloat(totalAmout),
		Details:        details,
	}
}

func (t *TransactionEntity) ToResponse() *TransactionResponse {
	if len(t.Details) == 0 {
		return nil
	}

	details := make([]*TransactionDetailResponse, len(t.Details))
	for i, detail := range t.Details {
		details[i] = &TransactionDetailResponse{
			ID:          detail.ID,
			ProductID:   detail.ProductID,
			ProductName: detail.ProductName,
			Quantity:    detail.Quantity,
			Price:       detail.Price.InexactFloat64(),
		}
	}

	return &TransactionResponse{
		ID:             t.ID,
		IdempotencyKey: t.IdempotencyKey,
		CustomerID:     t.CustomerID,
		CartID:         t.CartID,
		Status:         t.Status.Enum(),
		TotalAmount:    t.TotalAmount.InexactFloat64(),
		PaymentMethod:  t.PaymentMethod.Enum(),
		Details:        details,
	}
}
