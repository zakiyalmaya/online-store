package category

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	mockCategoryRepo "github.com/zakiyalmaya/online-store/mocks/infrastructure/repository/category"
	"github.com/zakiyalmaya/online-store/model"
)

var (
	mockCaregoryRepository *mockCategoryRepo.MockRepository
	categorySvc      Service
)

func Setup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	
	mockCaregoryRepository = mockCategoryRepo.NewMockRepository(mockCtl)
	categorySvc = NewCategoryService(&repository.Repositories{
		Category: mockCaregoryRepository,
	})
}

func TestCreate(t *testing.T) {
	Setup(t)
	
	testCases := []struct {
		name     string
		categoryName string
		mock     func()
		wantErr  bool
	}{
		{
			name:     "Given valid request when create category then return success",
			categoryName: "Fashion",
			mock: func() {
				mockCaregoryRepository.EXPECT().Create(&model.CategoryEntity{Name: "Fashion"}).Return(nil)
			},
			wantErr: false,
		},
		{
			name:     "Given error when create category then return error",
			categoryName: "Fashion",
			mock: func() {
				mockCaregoryRepository.EXPECT().Create(&model.CategoryEntity{Name: "Fashion"}).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := categorySvc.Create(tc.categoryName)
			if (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	Setup(t)
	
	testCases := []struct {
		name     string
		mock     func()
		wantErr  bool
	}{
		{
			name: "Given valid request when get all category then return success",
			mock: func() {
				mockCaregoryRepository.EXPECT().GetAll().Return([]*model.CategoryEntity{{Name: "Fashion"}}, nil)
			},
			wantErr: false,
		},
		{
			name: "Given error when get all category then return error",
			mock: func() {
				mockCaregoryRepository.EXPECT().GetAll().Return(nil, errors.New("error"))
			},
			wantErr: true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			_, err := categorySvc.GetAll()
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
