package product

import "github.com/zakiyalmaya/online-store/model"

type Service interface {
	Create(request *model.ProductRequest) error
	GetAll(categoryID *int) ([]*model.ProductResponse, error)
}