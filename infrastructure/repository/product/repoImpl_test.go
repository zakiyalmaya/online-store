package product

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
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

	testCases := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when create then return success",
			mock: func() {
				mock.ExpectExec("INSERT INTO products (name, description, price, stock_quantity, category_id) VALUES (?, ?, ?, ?, ?)").
					WithArgs("T-Shirt", "T-Shirt description", "10000", 10, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Given error when create then return error",
			mock: func() {
				mock.ExpectExec("INSERT INTO products (name, description, price, stock_quantity, category_id) VALUES (?, ?, ?, ?, ?)").
					WithArgs("T-Shirt", "T-Shirt description", "10000", 10, 1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewProductRepository(sqlxDB)
			tc.mock()
			err := repo.Create(&model.ProductEntity{
				Name:          "T-Shirt",
				Price:         decimal.NewFromFloat(10000),
				StockQuantity: 10,
				Description:   "T-Shirt description",
				CategoryID:    1,
			})
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
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
			name: "Given valid request when get by id then return success",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, description, price, stock_quantity, category_id, created_at, updated_at FROM products WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "stock_quantity", "category_id", "created_at", "updated_at"}).
						AddRow(1, "T-Shirt", "T-Shirt description", "10000", 10, 1, time.Time{}, time.Time{}))
			},
			wantErr: false,
		},
		{
			name: "Given error when get by id then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, description, price, stock_quantity, category_id, created_at, updated_at FROM products WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewProductRepository(sqlxDB)
			tc.mock()
			_, err := repo.GetByID(1)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	categoryID := 1
	request := &model.GetProductRequest{
		CategoryID: &categoryID,
		Limit:      10,
		Page:       1,
	}

	testCases := []struct {
		name    string
		request *model.GetProductRequest
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when get all then return success",
			request: request,
			mock: func() {
				mock.ExpectQuery("SELECT p.id, p.name, p.description, p.price, p.stock_quantity, c.name FROM products AS p JOIN categories AS c ON p.category_id = c.id WHERE TRUE AND p.category_id = ? LIMIT ? OFFSET ?").
					WithArgs(request.CategoryID, request.Limit, (request.Page-1)*request.Limit).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "stock_quantity", "name"}).
						AddRow(1, "T-Shirt", "T-Shirt description", "10000", 10, "Fashion"))
			},
			wantErr: false,
		},
		{
			name: "Given error when get all then return error",
			request: request,
			mock: func() {
				mock.ExpectQuery("SELECT p.id, p.name, p.description, p.price, p.stock_quantity, c.name FROM products AS p JOIN categories AS c ON p.category_id = c.id WHERE TRUE AND p.category_id = ? LIMIT ? OFFSET ?").
					WithArgs(request.CategoryID, request.Limit, (request.Page-1)*request.Limit).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given error scan row when get all then return error",
			request: request,
			mock:    func() {
				mock.ExpectQuery("SELECT p.id, p.name, p.description, p.price, p.stock_quantity, c.name FROM products AS p JOIN categories AS c ON p.category_id = c.id WHERE TRUE AND p.category_id = ? LIMIT ? OFFSET ?").
					WithArgs(request.CategoryID, request.Limit, (request.Page-1)*request.Limit).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description", "price", "stock_quantity", "name", "id"}).
						AddRow(1, "T-Shirt", "T-Shirt description", "10000", 10, "Fashion", 1))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewProductRepository(sqlxDB)
			tc.mock()
			_, err := repo.GetAll(tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
