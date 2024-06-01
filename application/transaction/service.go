package transaction

import "github.com/zakiyalmaya/online-store/model"


type Service interface {
	Checkout(request *model.TransactionRequest) (*model.TransactionResponse, error)
	GetByID(id int) (*model.TransactionResponse, error)
}