package cart

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
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

	items := []*model.CartItemEntity{
		{
			ProductID: 1,
			Quantity:  1,
		},
	}

	request := &model.CartEntity{
		CustomerID: 1,
		Status:     cartEnum.CartStatusActive,
		Items:      items,
	}

	testCases := []struct {
		name    string
		request *model.CartEntity
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when create then return success",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO shopping_carts (customer_id, status) VALUES (?, ?)").
					WithArgs(request.CustomerID, request.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO cart_items (shopping_cart_id, product_id, quantity) VALUES (?, ?, ?)").
					WithArgs(1, request.Items[0].ProductID, request.Items[0].Quantity).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "shopping_cart_id", "product_id", "quantity", "price", "product_name"}).
						AddRow(1, 1, 1, 1, 1000, "Product 1"))
			},
			wantErr: false,
		},
		{
			name:    "Given error begin db transaction when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin().WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given error insert shopping carts when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO shopping_carts (customer_id, status) VALUES (?, ?)").
					WithArgs(request.CustomerID, request.Status).
					WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Given error insert cart items when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO shopping_carts (customer_id, status) VALUES (?, ?)").
					WithArgs(request.CustomerID, request.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO cart_items (shopping_cart_id, product_id, quantity) VALUES (?, ?, ?)").
					WithArgs(1, request.Items[0].ProductID, request.Items[0].Quantity).
					WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{

			name:    "Given error getting cart items when create then return success",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO shopping_carts (customer_id, status) VALUES (?, ?)").
					WithArgs(request.CustomerID, request.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO cart_items (shopping_cart_id, product_id, quantity) VALUES (?, ?, ?)").
					WithArgs(1, request.Items[0].ProductID, request.Items[0].Quantity).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given error getting carts when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO shopping_carts (customer_id, status) VALUES (?, ?)").
					WithArgs(request.CustomerID, request.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO cart_items (shopping_cart_id, product_id, quantity) VALUES (?, ?, ?)").
					WithArgs(1, request.Items[0].ProductID, request.Items[0].Quantity).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name:    "Given error commit when create then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO shopping_carts (customer_id, status) VALUES (?, ?)").
					WithArgs(request.CustomerID, request.Status).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec("INSERT INTO cart_items (shopping_cart_id, product_id, quantity) VALUES (?, ?, ?)").
					WithArgs(1, request.Items[0].ProductID, request.Items[0].Quantity).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit().WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCartRepository(sqlxDB)
			tc.mock()
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
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "shopping_cart_id", "product_id", "quantity", "price", "product_name"}).
						AddRow(1, 1, 1, 1, 1000, "Product 1"))
			},
			wantErr: false,
		},
		{
			name: "Given error getting shopping cart when get by id then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "Given error getting cart items when get by id then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCartRepository(sqlxDB)
			tc.mock()
			_, err := repo.GetByID(1)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestGetByParams(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	status := int(cartEnum.CartStatusActive)

	testCases := []struct {
		name    string
		request *model.GetCartRequest
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when get by param then return success",
			request: &model.GetCartRequest{
				CustomerID: 1,
				Status:     &status,
			},
			mock: func() {
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE TRUE AND customer_id = ? AND status = ? ORDER BY id").
					WithArgs(1, status).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id DESC").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "shopping_cart_id", "product_id", "quantity", "price", "product_name"}).
						AddRow(1, 1, 1, 1, 1000, "Product 1"))
			},
			wantErr: false,
		},
		{
			name: "Given error getting cart items when get by params then return error",
			request: &model.GetCartRequest{
				CustomerID: 1,
				Status:     &status,
			},
			mock: func() {
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE TRUE AND customer_id = ? AND status = ? ORDER BY id").
					WithArgs(1, status).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id DESC").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "Given error getting shopping cart when get by params then return error",
			request: &model.GetCartRequest{
				CustomerID: 1,
				Status:     &status,
			},
			mock: func() {
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE TRUE AND customer_id = ? AND status = ? ORDER BY id").
					WithArgs(1, status).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "Given empty shopping cart when get by params then return error",
			request: &model.GetCartRequest{
				CustomerID: 1,
				Status:     &status,
			},
			mock: func() {
				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE TRUE AND customer_id = ? AND status = ? ORDER BY id").
					WithArgs(1, status).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}))
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCartRepository(sqlxDB)
			tc.mock()
			_, err := repo.GetByParams(tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetByParams() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestUpsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	request := []*model.CartItemEntity{
		{
			ProductID: 1,
			Quantity:  1,
			CartID:    1,
		},
	}

	testCases := []struct {
		name    string
		request []*model.CartItemEntity
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when upsert then return success",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO cart_items (product_id, shopping_cart_id, quantity) VALUES (?, ?, ?) ON CONFLICT(product_id, shopping_cart_id) DO UPDATE SET quantity = cart_items.quantity + excluded.quantity").
					WithArgs(request[0].ProductID, request[0].CartID, request[0].Quantity).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

				mock.ExpectQuery("SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?").
					WithArgs(request[0].CartID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "customer_id", "status", "created_at", "updated_at"}).
						AddRow(1, 1, cartEnum.CartStatusActive, time.Time{}, time.Time{}))

				mock.ExpectQuery("SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ? ORDER BY ci.id").
					WithArgs(request[0].CartID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "shopping_cart_id", "product_id", "quantity", "price", "product_name"}).
						AddRow(1, 1, 1, 1, 1000, "Product 1"))
			},
			wantErr: false,
		},
		{
			name:    "Given error when upsert then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO cart_items (product_id, shopping_cart_id, quantity) VALUES (?, ?, ?) ON CONFLICT(product_id, shopping_cart_id) DO UPDATE SET quantity = cart_items.quantity + excluded.quantity").
					WithArgs(request[0].ProductID, request[0].CartID, request[0].Quantity).
					WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Given error commit db transaction when upsert then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO cart_items (product_id, shopping_cart_id, quantity) VALUES (?, ?, ?) ON CONFLICT(product_id, shopping_cart_id) DO UPDATE SET quantity = cart_items.quantity + excluded.quantity").
					WithArgs(request[0].ProductID, request[0].CartID, request[0].Quantity).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit().WillReturnError(errors.New("error"))
					
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name:    "Given error begin db transaction when upsert then return error",
			request: request,
			mock: func() {
				mock.ExpectBegin().WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCartRepository(sqlxDB)
			tc.mock()
			_, err := repo.Upsert(1, tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("Upsert() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestDelete(t *testing.T) {
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
			name: "Given valid request when delete then return success",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM cart_items WHERE id = ? AND EXISTS (SELECT 1 FROM shopping_carts WHERE id = ? AND status = ? AND customer_id = ?)").
					WithArgs(1, 1, cartEnum.CartStatusActive, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Given error when delete then return error",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM cart_items WHERE id = ? AND EXISTS (SELECT 1 FROM shopping_carts WHERE id = ? AND status = ? AND customer_id = ?)").
					WithArgs(1, 1, cartEnum.CartStatusActive, 1).
					WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "Given error commit db transaction when delete then return error",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM cart_items WHERE id = ? AND EXISTS (SELECT 1 FROM shopping_carts WHERE id = ? AND status = ? AND customer_id = ?)").
					WithArgs(1, 1, cartEnum.CartStatusActive, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit().WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "Given error begin db transaction when delete then return error",
			mock: func() {
				mock.ExpectBegin().WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "Given error affected rows when delete then return error",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("DELETE FROM cart_items WHERE id = ? AND EXISTS (SELECT 1 FROM shopping_carts WHERE id = ? AND status = ? AND customer_id = ?)").
					WithArgs(1, 1, cartEnum.CartStatusActive, 1).
					WillReturnResult(sqlmock.NewResult(0, 0))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCartRepository(sqlxDB)
			tc.mock()
			err := repo.Delete(&model.DeleteCartRequest{
				CartItemID: 1,
				ID:     1,
				CustomerID: 1,
				Status: cartEnum.CartStatusActive,
			})
			if (err != nil) != tc.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestGetItemByID(t *testing.T) {
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
			name: "Given valid request when get item by id then return success",
			mock: func() {
				mock.ExpectQuery("SELECT id, shopping_cart_id, product_id, quantity FROM cart_items WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "shopping_cart_id", "product_id", "quantity"}).
						AddRow(1, 1, 1, 1))
			},
			wantErr: false,
		},
		{
			name: "Given error when get item by id then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, shopping_cart_id, product_id, quantity FROM cart_items WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCartRepository(sqlxDB)
			tc.mock()
			_, err := repo.GetItemByID(1)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetItemByID() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
