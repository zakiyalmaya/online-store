package category

import (
	"fmt"

	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/model"
)

type categorySvcImpl struct {
	repos *repository.Repositories
}

func NewCategoryService(repos *repository.Repositories) Service {
	return &categorySvcImpl{repos: repos}
}

func (c *categorySvcImpl) Create(name string) error {
	err := c.repos.Category.Create(&model.CategoryEntity{Name: name})
	if err != nil {
		return fmt.Errorf("error creating category")
	}

	return nil
}

func (c *categorySvcImpl) GetAll() ([]*model.CategoryEntity, error) {
	categories, err := c.repos.Category.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting all categories")
	}
	
	return categories, nil
}