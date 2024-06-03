package category

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=CategoryRepository.go
type Repository interface {
	Create(category *model.CategoryEntity) error
	GetAll() ([]*model.CategoryEntity, error)
	GetByID(id int) (*model.CategoryEntity, error)
}
