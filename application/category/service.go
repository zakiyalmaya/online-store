package category

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=CategoryService.go
type Service interface {
	Create(name string) error
	GetAll() ([]*model.CategoryEntity, error)
}
