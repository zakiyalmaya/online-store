package category

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/online-store/model"
)

type categoryRepoImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepoImpl{db: db}
}

func (c *categoryRepoImpl) Create(category *model.CategoryEntity) error {
	query := "INSERT INTO categories (name) VALUES (:name)"
	_, err := c.db.NamedExec(query, category)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}
	
	return nil
}

func (c *categoryRepoImpl) GetAll() ([]*model.CategoryEntity, error) {
	categories := make([]*model.CategoryEntity, 0)
	query := "SELECT id, name, created_at, updated_at FROM categories"
	res, err := c.db.Query(query)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for res.Next() {
		category := &model.CategoryEntity{}
		if err := res.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}