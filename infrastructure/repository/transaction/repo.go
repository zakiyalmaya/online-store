package transaction

import "github.com/zakiyalmaya/online-store/model"

type Repository interface {
	Create(transaction *model.TransactionEntity) (*model.TransactionEntity, error)
	GetByID(id int) (*model.TransactionEntity, error)
}