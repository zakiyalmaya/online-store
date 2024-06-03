package product

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=ProductRepository.go
type Repository interface {
	Create(product *model.ProductEntity) error
	GetAll(request *model.GetProductRequest) ([]*model.ProductResponse, error)
	GetByID(id int) (*model.ProductEntity, error)
}