package cart

import (
	"github.com/zakiyalmaya/online-store/model"
)

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=CartService.go
type Service interface {
	Create(request *model.CreateCartRequest) (*model.CartResponse, error)
	GetByParams(request *model.GetCartRequest) ([]*model.CartResponse, error)
	Delete(request *model.DeleteCartRequest) error
}