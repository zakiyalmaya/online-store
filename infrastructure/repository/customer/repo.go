package customer

import "github.com/zakiyalmaya/online-store/model"

type CustomerRepository interface {
	Create(customer *model.CustomerEntity) error
	GetByUsername(username string) (*model.CustomerEntity, error)
}
