package transaction

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	transactionEnum "github.com/zakiyalmaya/online-store/constant/transaction"
	cartEnum "github.com/zakiyalmaya/online-store/constant/cart"
	"github.com/zakiyalmaya/online-store/model"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	request := &model.TransactionEntity{
		TotalAmount:     decimal.NewFromFloat(10000),
		Status:    transactionEnum.TransactionStatusInprogress,
		CustomerID: 1,
		Details: []*model.TransactionDetailEntity{
			{
				ProductID: 1,
				Quantity:  1,
				Price:     decimal.NewFromFloat(10000),
			},
		},
	}

	testCases := []struct {
		name    string
		request *model.TransactionEntity
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when create then return success",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (?, ?, ?, ?)").
					WithArgs(1, request.Details[0].ProductID, request.Details[0].Quantity, request.Details[0].Price).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE shopping_carts SET status = ? WHERE id = ?").
					WithArgs(cartEnum.CartStatusPending, request.CartID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "idempotency_key", "customer_id", "shopping_cart_id", "status", "total_amount", "payment_method", "created_at", "updated_at"}).
						AddRow(1, request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT td.id, td.transaction_id, td.product_id, p.name AS product_name, td.quantity, td.price, td.created_at, td.updated_at FROM transaction_details AS td JOIN products AS p ON td.product_id = p.id WHERE td.transaction_id = ? ORDER BY td.id").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "transaction_id", "product_id", "product_name", "quantity", "price", "created_at", "updated_at"}).
						AddRow(1, 1, request.Details[0].ProductID, request.Details[0].ProductName, request.Details[0].Quantity, request.Details[0].Price, time.Time{}, time.Time{}))
			},
			wantErr: false,
		},
		{
			name:    "Given error begin db transaction when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin().WillReturnError(errors.New("error begin db transaction"))
			},
			wantErr: true,
		},
		{
			name:    "Given error insert transaction when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnError(errors.New("error insert transaction"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Given error insert transaction details when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (?, ?, ?, ?)").
					WithArgs(1, request.Details[0].ProductID, request.Details[0].Quantity, request.Details[0].Price).
					WillReturnError(errors.New("error insert transaction details"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Given error update shopping cart when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (?, ?, ?, ?)").
					WithArgs(1, request.Details[0].ProductID, request.Details[0].Quantity, request.Details[0].Price).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE shopping_carts SET status = ? WHERE id = ?").
					WithArgs(cartEnum.CartStatusPending, request.CartID).
					WillReturnError(errors.New("error update shopping cart"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Given error committing transaction when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (?, ?, ?, ?)").
					WithArgs(1, request.Details[0].ProductID, request.Details[0].Quantity, request.Details[0].Price).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE shopping_carts SET status = ? WHERE id = ?").
					WithArgs(cartEnum.CartStatusPending, request.CartID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit().WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given error getting transaction when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (?, ?, ?, ?)").
					WithArgs(1, request.Details[0].ProductID, request.Details[0].Quantity, request.Details[0].Price).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE shopping_carts SET status = ? WHERE id = ?").
					WithArgs(cartEnum.CartStatusPending, request.CartID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},		
		{
			name:    "Given error getting transaction details when create then return error",
			request: &model.TransactionEntity{},
			mock:    func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs(request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (?, ?, ?, ?)").
					WithArgs(1, request.Details[0].ProductID, request.Details[0].Quantity, request.Details[0].Price).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("UPDATE shopping_carts SET status = ? WHERE id = ?").
					WithArgs(cartEnum.CartStatusPending, request.CartID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "idempotency_key", "customer_id", "shopping_cart_id", "status", "total_amount", "payment_method", "created_at", "updated_at"}).
						AddRow(1, request.IdempotencyKey, request.CustomerID, request.CartID, request.Status, request.TotalAmount, request.PaymentMethod, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT td.id, td.transaction_id, td.product_id, p.name AS product_name, td.quantity, td.price, td.created_at, td.updated_at FROM transaction_details AS td JOIN products AS p ON td.product_id = p.id WHERE td.transaction_id = ? ORDER BY td.id").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},	
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			repo := NewTransactionRepository(sqlxDB)
			_, err := repo.Create(tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	testCases := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid id when get by id then return success",
			mock: func() {
				mock.ExpectQuery("SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "idempotency_key", "customer_id", "shopping_cart_id", "status", "total_amount", "payment_method", "created_at", "updated_at"}).
						AddRow(1, "idempotency_key", 1, 1, 1, "1000", 1, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT td.id, td.transaction_id, td.product_id, p.name AS product_name, td.quantity, td.price, td.created_at, td.updated_at FROM transaction_details AS td JOIN products AS p ON td.product_id = p.id WHERE td.transaction_id = ? ORDER BY td.id").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "transaction_id", "product_id", "product_name", "quantity", "price", "created_at", "updated_at"}).
						AddRow(1, 1, 1, "product_name", 1, "1000", time.Time{}, time.Time{}))
			},
			wantErr: false,
		},
		{
			name: "Given error getting transaction details when get by id then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "idempotency_key", "customer_id", "shopping_cart_id", "status", "total_amount", "payment_method", "created_at", "updated_at"}).
						AddRow(1, "idempotency_key", 1, 1, 1, "1000", 1, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT td.id, td.transaction_id, td.product_id, p.name AS product_name, td.quantity, td.price, td.created_at, td.updated_at FROM transaction_details AS td JOIN products AS p ON td.product_id = p.id WHERE td.transaction_id = ? ORDER BY td.id").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "Given error getting transaction when get by id then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			repo := NewTransactionRepository(sqlxDB)
			_, err := repo.GetByID(1)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
