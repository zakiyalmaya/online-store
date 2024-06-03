package transaction

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=TransactionService.go
type Service interface {
	Checkout(request *model.TransactionRequest) (*model.TransactionResponse, error)
	GetByID(id int) (*model.TransactionResponse, error)
}