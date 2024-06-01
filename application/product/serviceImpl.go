package product

import (
	"database/sql"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/model"
)

type productSvcImpl struct {
	repos *repository.Repositories
}

func NewProductService(repos *repository.Repositories) Service {
	return &productSvcImpl{repos: repos}
}

func (p *productSvcImpl) Create(request *model.ProductRequest) error {
	category, err := p.repos.Category.GetByID(request.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("category not found")
		}

		return fmt.Errorf("error getting category by id")
	}

	if err := p.repos.Product.Create(&model.ProductEntity{
		Name:          request.Name,
		Description:   request.Description,
		Price:         decimal.NewFromFloat(request.Price),
		CategoryID:    category.ID,
		StockQuantity: request.StockQuantity,
	}); err != nil {
		return fmt.Errorf("error creating product")
	}

	return nil
}

func (p *productSvcImpl) GetAll(categoryID *int) ([]*model.ProductResponse, error) {
	products, err := p.repos.Product.GetAll(categoryID)
	if err != nil {
		return nil, fmt.Errorf("error getting all products")
	}

	return products, nil
}
