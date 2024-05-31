package customer

import "github.com/zakiyalmaya/online-store/model"

type Service interface {
	Register(request *model.CustomerRequest) error
	Login(request *model.AuthRequest) (*model.AuthResponse, error)
	Logout(username string) error
}