package product

import "github.com/zakiyalmaya/online-store/model"

type Service interface {
	Create(request *model.CreateProductRequest) error
	GetAll(request *model.GetProductRequest) ([]*model.ProductResponse, error)
}