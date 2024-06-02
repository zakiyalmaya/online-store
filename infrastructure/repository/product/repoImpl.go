package product

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/online-store/model"
)

type productRepoImpl struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) Repository {
	return &productRepoImpl{db: db}
}

func (p *productRepoImpl) Create(product *model.ProductEntity) error {
	query := "INSERT INTO products (name, description, price, stock_quantity, category_id) VALUES (:name, :description, :price, :stock_quantity, :category_id)"
	_, err := p.db.NamedExec(query, product)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (p *productRepoImpl) GetAll(request *model.GetProductRequest) ([]*model.ProductResponse, error) {
	products := make([]*model.ProductResponse, 0)
	params := make([]interface{}, 0)
	
	query := "SELECT p.id, p.name, p.description, p.price, p.stock_quantity, c.name FROM products AS p JOIN categories AS c ON p.category_id = c.id WHERE TRUE"
	if request.CategoryID != nil {
		query += " AND p.category_id = ?"
		params = append(params, request.CategoryID)
	}

	if request.Limit != 0 {
		query += " LIMIT ?"
		params = append(params, request.Limit)
	}

	if request.Page != 0 {
		query += " OFFSET ?"
		params = append(params, (request.Page-1)*request.Limit)
	}

	res, err := p.db.Query(query, params...)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for res.Next() {
		product := &model.ProductResponse{}
		if err := res.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.StockQuantity,
			&product.Category,
		); err != nil {
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
		products = append(products, product)
	}	

	return products, nil
}

func (p *productRepoImpl) GetByID(id int) (*model.ProductEntity, error) {
	product := &model.ProductEntity{}
	query := "SELECT id, name, description, price, stock_quantity, category_id, created_at, updated_at FROM products WHERE id = ?"

	err := p.db.Get(product, query, id)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return product, nil
}