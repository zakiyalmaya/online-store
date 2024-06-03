package cart

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	cartEnum "github.com/zakiyalmaya/online-store/constant/cart"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	mockCartRepo "github.com/zakiyalmaya/online-store/mocks/infrastructure/repository/cart"
	mockProductRepo "github.com/zakiyalmaya/online-store/mocks/infrastructure/repository/product"
	"github.com/zakiyalmaya/online-store/model"
)

var (
	mockCartRepository    *mockCartRepo.MockRepository
	mockProductRepository *mockProductRepo.MockRepository
	cartSvc               Service
)

func Setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCartRepository = mockCartRepo.NewMockRepository(ctrl)
	cartSvc = NewCartService(&repository.Repositories{
		Cart:    mockCartRepository,
		Product: mockProductRepository,
	})
}

func TestGetByParams(t *testing.T) {
	Setup(t)

	status := int(cartEnum.CartStatusActive)
	request := &model.GetCartRequest{
		CustomerID: 1,
		Status:     &status,
	}

	testCases := []struct {
		name    string
		request *model.GetCartRequest
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when get cart by params then return success",
			request: request,
			mock: func() {
				mockCartRepository.EXPECT().GetByParams(request).Return([]*model.CartEntity{
					{
						ID:         1,
						CustomerID: 1,
						Status:     cartEnum.CartStatusActive,
						Items: []*model.CartItemEntity{
							{
								ProductID: 1,
								Quantity:  1,
								Price:     decimal.NewFromFloat(10000),
							},
						},
					},
				}, nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:    "Given error when get cart by params then return error",
			request: request,
			mock: func() {
				mockCartRepository.EXPECT().GetByParams(request).Return(nil, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			_, err := cartSvc.GetByParams(tc.request)
			if err != nil && !tc.wantErr {
				t.Errorf("GetByParams() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func TestDelete(t *testing.T) {
	Setup(t)

	request := &model.DeleteCartRequest{
		ID:         1,
		CartItemID: 1,
		CustomerID: 1,
		Status:     cartEnum.CartStatusActive,
	}

	testCases := []struct {
		name    string
		request *model.DeleteCartRequest
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when delete cart then return success",
			request: request,
			mock: func() {
				mockCartRepository.EXPECT().GetItemByID(request.CartItemID).Return(&model.CartItemEntity{
					ID:     1,
					CartID: 1,
				}, nil).Times(1)

				mockCartRepository.EXPECT().Delete(request).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:    "Given error when delete cart then return error",
			request: request,
			mock: func() {
				mockCartRepository.EXPECT().GetItemByID(request.CartItemID).Return(&model.CartItemEntity{
					ID:     1,
					CartID: 1,
				}, nil).Times(1)

				mockCartRepository.EXPECT().Delete(request).Return(errors.New("error")).Times(1)
			},
			wantErr: true,
		},
		{
			name:    "Given error when get cart item by id then return error",
			request: request,
			mock: func() {
				mockCartRepository.EXPECT().GetItemByID(request.CartItemID).Return(nil, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := cartSvc.Delete(tc.request)
			if err != nil && !tc.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
