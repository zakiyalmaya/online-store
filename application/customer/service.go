package customer

import "github.com/zakiyalmaya/online-store/model"

//go:generate go run github.com/golang/mock/mockgen --build_flags=--mod=vendor -package mocks -source=service.go -destination=CustomerService.go
type Service interface {
	Register(request *model.CustomerRequest) error
	Login(request *model.AuthRequest) (*model.AuthResponse, error)
	Logout(username string) error
}