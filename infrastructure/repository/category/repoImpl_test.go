package category

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
		categoryName string
		mock     func()
		wantErr  bool
	}{
		{
			name:     "Given valid request when create then return success",
			categoryName: "Fashion",
			mock: func() {
				mock.ExpectExec("INSERT INTO categories (name) VALUES (?)").
					WithArgs("Fashion").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:     "Given error when create then return error",
			categoryName: "Fashion",
			mock: func() {
				mock.ExpectExec("INSERT INTO categories (name) VALUES (?)").
					WithArgs("Fashion").
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
	},
}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCategoryRepository(sqlxDB)
			tc.mock()
			err := c.Create(&model.CategoryEntity{
				Name: tc.categoryName,
			})
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
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

	testCases := []struct {
		name     string
		mock     func()
		wantErr  bool
	}{
		{
			name: "Given valid request when getting all categories then return success",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, created_at, updated_at FROM categories").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
						AddRow(1, "Fashion", time.Time{}, time.Time{}))
			},
			wantErr: false,
		},
		{
			name: "Given error when getting all categories then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, created_at, updated_at FROM categories").
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "Given error scan row when getting all categories then return error",
			mock: func() {
				mock.ExpectQuery("SELECT id, name, created_at, updated_at FROM categories").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
						AddRow(1, "Fashion"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCategoryRepository(sqlxDB)
			tc.mock()
			_, err := c.GetAll()
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tc.wantErr)
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
		name     string
		categoryID int
		mock     func()
		wantErr  bool
	}{
		{
			name:     "Given valid id when get by id then return success",
			categoryID: 1,
			mock: func() {
				mock.ExpectQuery("SELECT id, name, created_at, updated_at FROM categories WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
						AddRow(1, "Fashion", time.Time{}, time.Time{}))
			},
			wantErr: false,
		},
		{
			name:     "Given error when get by id then return error",
			categoryID: 1,
			mock: func() {
				mock.ExpectQuery("SELECT id, name, created_at, updated_at FROM categories WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCategoryRepository(sqlxDB)
			tc.mock()
			_, err := c.GetByID(tc.categoryID)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}