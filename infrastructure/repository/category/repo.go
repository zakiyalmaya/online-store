package category

import "github.com/zakiyalmaya/online-store/model"

type CategoryRepository interface {
	Create(category *model.CategoryEntity) error
	GetAll() ([]*model.CategoryEntity, error)
}
