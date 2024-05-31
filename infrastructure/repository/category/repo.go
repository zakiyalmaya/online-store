package category

import "github.com/zakiyalmaya/online-store/model"

type Repository interface {
	Create(category *model.CategoryEntity) error
	GetAll() ([]*model.CategoryEntity, error)
	GetByID(id int) (*model.CategoryEntity, error)
}
