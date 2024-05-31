package customer

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/online-store/model"
)

type customerRepoImpl struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepoImpl{db: db}
}

func (c *customerRepoImpl) Create(customer *model.CustomerEntity) error {
	query := "INSERT INTO customers (name, username, password, email, phone_number, address) VALUES (:name, :username, :password, :email, :phone_number, :address)"
	_, err := c.db.NamedExec(query, customer)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (c *customerRepoImpl) GetByUsername(username string) (*model.CustomerEntity, error) {
	customer := &model.CustomerEntity{}
	query := "SELECT id, name, username, password, phone_number, email, address, created_at, updated_at FROM customers WHERE username = ?"

	err := c.db.Get(customer, query, username)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return customer, nil
}
