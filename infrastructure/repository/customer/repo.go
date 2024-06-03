package customer

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=repo.go -destination=CustomerRepository.go
type Repository interface {
	Create(customer *model.CustomerEntity) error
	GetByUsername(username string) (*model.CustomerEntity, error)
}
