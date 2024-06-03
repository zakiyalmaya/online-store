package customer

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
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
		name     string
		customer model.CustomerEntity
		mock     func()
		wantErr  bool
	}{
		{
			name: "Given valid request when create then return success",
			customer: model.CustomerEntity{
				Name:        "John",
				Email:       "KUZuL@example.com",
				Username:    "john",
				Password:    "John-123",
				PhoneNumber: "08123456789",
				Address:     "Jl. Raya",
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO customers (name, username, password, email, phone_number, address) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs("John", "john", "John-123", "KUZuL@example.com", "08123456789", "Jl. Raya").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Given error when create then return error",
			customer: model.CustomerEntity{
				Name:        "John",
				Email:       "KUZuL@example.com",
				Username:    "john",
				Password:    "John-123",
				PhoneNumber: "08123456789",
				Address:     "Jl. Raya",
			},
			mock: func() {
				mock.ExpectExec("INSERT INTO customers (name, username, password, email, phone_number, address) VALUES (?, ?, ?, ?, ?, ?)").
					WithArgs("John", "john", "John-123", "KUZuL@example.com", "08123456789", "Jl. Raya").
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			repo := NewCustomerRepository(sqlxDB)
			err := repo.Create(&tc.customer)
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	testCases := []struct {
		name     string
		username string
		mock     func()
		wantErr  bool
	}{
		{
			name:     "Given valid request when get by username then return success",
			username: "john",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, username, password, phone_number, email, address, created_at, updated_at FROM customers WHERE username = ?").
					WithArgs("john").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "username", "password", "email", "phone_number", "address", "created_at", "updated_at"}).
						AddRow(1, "John", "john", "John-123", "KUZuL@example.com", "08123456789", "Jl. Raya", time.Time{}, time.Time{}))
			},
			wantErr: false,
		},
		{
			name:     "Given error when get by username then return error",
			username: "john",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, username, password, phone_number, email, address, created_at, updated_at FROM customers WHERE username = ?").
					WithArgs("john").
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			repo := NewCustomerRepository(sqlxDB)
			_, err := repo.GetByUsername(tc.username)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
