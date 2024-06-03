package product

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	mockCategoryRepo "github.com/zakiyalmaya/online-store/mocks/infrastructure/repository/category"
	mockProductRepo "github.com/zakiyalmaya/online-store/mocks/infrastructure/repository/product"
	"github.com/zakiyalmaya/online-store/model"
)

var (
	mockProductRepository  *mockProductRepo.MockRepository
	mockCategoryRepository *mockCategoryRepo.MockRepository
	productSvc             Service
)

func Setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository = mockProductRepo.NewMockRepository(ctrl)
	mockCategoryRepository = mockCategoryRepo.NewMockRepository(ctrl)

	repos := &repository.Repositories{
		Product:  mockProductRepository,
		Category: mockCategoryRepository,
	}
	productSvc = NewProductService(repos)
}

func TestCreate(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid request when create product then return success",
			mock: func() {
				mockCategoryRepository.EXPECT().GetByID(1).Return(&model.CategoryEntity{
					ID: 1,
				}, nil).Times(1)

				mockProductRepository.EXPECT().Create(&model.ProductEntity{
					Name:          "T-Shirt",
					Description:   "T-Shirt description",
					Price:         decimal.NewFromFloat(10000),
					StockQuantity: 10,
					CategoryID:    1,
				}).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name: "Given error when create product then return error",
			mock: func() {
				mockCategoryRepository.EXPECT().GetByID(1).Return(&model.CategoryEntity{
					ID: 1,
				}, nil).Times(1)

				mockProductRepository.EXPECT().Create(&model.ProductEntity{
					Name:          "T-Shirt",
					Description:   "T-Shirt description",
					Price:         decimal.NewFromFloat(10000),
					StockQuantity: 10,
					CategoryID:    1,
				}).Return(errors.New("error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Given error when get category by id then return error",
			mock: func() {
				mockCategoryRepository.EXPECT().GetByID(1).Return(nil, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "Given error no rows when get category by id then return error",
			mock: func() {
				mockCategoryRepository.EXPECT().GetByID(1).Return(nil, sql.ErrNoRows).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := productSvc.Create(&model.CreateProductRequest{
				Name:          "T-Shirt",
				Description:   "T-Shirt description",
				Price:         10000,
				StockQuantity: 10,
				CategoryID:    1,
			})
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	Setup(t)

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
			name:    "Given valid request when get all product then return success",
			request: request,
			mock: func() {
				mockProductRepository.EXPECT().GetAll(request).Return([]*model.ProductResponse{
					{
						ID:            1,
						Name:          "T-Shirt",
						Description:   "T-Shirt description",
						Price:         10000,
						StockQuantity: 10,
						Category:      "Fashion",
					},
				}, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:    "Given error when get all product then return error",
			request: request,
			mock: func() {
				mockProductRepository.EXPECT().GetAll(request).Return(nil, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			_, err := productSvc.GetAll(tc.request)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
