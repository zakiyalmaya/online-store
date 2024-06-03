package product

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=ProductService.go
type Service interface {
	Create(request *model.CreateProductRequest) error
	GetAll(request *model.GetProductRequest) ([]*model.ProductResponse, error)
}