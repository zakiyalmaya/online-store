package product

import "github.com/zakiyalmaya/online-store/model"

type Repository interface {
	Create(product *model.ProductEntity) error
	GetAll(categoryID *int) ([]*model.ProductResponse, error)
	GetByID(id int) (*model.ProductEntity, error)
}