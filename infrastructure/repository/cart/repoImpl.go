package cart

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/online-store/model"
)

type cartRepoImpl struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) Repository {
	return &cartRepoImpl{db: db}
}

func (c *cartRepoImpl) Create(cart *model.CartEntity) (*model.CartEntity, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	res, err := tx.NamedExec(`INSERT INTO shopping_carts (customer_id, status) VALUES (:customer_id, :status)`, cart)
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	cartID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for _, item := range cart.Items {
		item.CartID = int(cartID)
		_, err = tx.NamedExec(`INSERT INTO cart_items (shopping_cart_id, product_id, quantity) VALUES (:shopping_cart_id, :product_id, :quantity)`, item)
		if err != nil {
			tx.Rollback()
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return c.getByID(int(cartID))
}

func (c *cartRepoImpl) getByID(id int) (*model.CartEntity, error) {
	cart := &model.CartEntity{}
	query := "SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE id = ?"
	err := c.db.Get(cart, query, id)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	items := []*model.CartItemEntity{}
	query = "SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ?"
	err = c.db.Select(&items, query, id)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	cart.Items = items
	return cart, nil
}

func (c *cartRepoImpl) GetByParams(request *model.GetCartRequest) ([]*model.CartEntity, error) {
	carts := []*model.CartEntity{}
	params := make([]interface{}, 0)
	query := "SELECT id, customer_id, status, created_at, updated_at FROM shopping_carts WHERE TRUE"

	if request.CustomerID != 0 {
		query += " AND customer_id = ?"
		params = append(params, request.CustomerID)
	}

	if request.Status != nil {
		query += " AND status = ?"
		params = append(params, request.Status)
	}

	res, err := c.db.Queryx(query, params...)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for res.Next() {
		cartEntity := &model.CartEntity{}
		if err := res.StructScan(cartEntity); err != nil {
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
		carts = append(carts, cartEntity)
	}

	if len(carts) == 0 {
		return nil, nil
	}

	for _, cart := range carts {
		items := []*model.CartItemEntity{}
		query = "SELECT ci.id, ci.shopping_cart_id, ci.product_id, ci.quantity, p.price, p.name AS product_name FROM cart_items AS ci JOIN products AS p ON ci.product_id = p.id WHERE shopping_cart_id = ?"
		res, err = c.db.Queryx(query, cart.ID)
		if err != nil {
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
		for res.Next() {
			item := &model.CartItemEntity{}
			if err := res.StructScan(item); err != nil {
				log.Println("errorRepository: ", err.Error())
				return nil, err
			}
			items = append(items, item)
		}
		cart.Items = items
	}

	return carts, nil
}

func (c *cartRepoImpl) Upsert(cartID int, items []*model.CartItemEntity) (*model.CartEntity, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for _, item := range items {
		queryUpsert := `
			INSERT INTO cart_items (product_id, shopping_cart_id, quantity)
			VALUES (:product_id, :shopping_cart_id, :quantity)
			ON CONFLICT(product_id, shopping_cart_id) DO UPDATE SET
				quantity = cart_items.quantity + excluded.quantity
		`
		_, err = tx.NamedExec(queryUpsert, item)
		if err != nil {
			tx.Rollback()
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return c.getByID(int(cartID))
}
