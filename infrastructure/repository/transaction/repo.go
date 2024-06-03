package transaction

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=TransactionRepository.go
type Repository interface {
	Create(transaction *model.TransactionEntity) (*model.TransactionEntity, error)
	GetByID(id int) (*model.TransactionEntity, error)
}