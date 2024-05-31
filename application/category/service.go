package category

import "github.com/zakiyalmaya/online-store/model"

type Service interface {
	Create(name string) error
	GetAll() ([]*model.CategoryEntity, error)
}
