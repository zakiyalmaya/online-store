package cart

import (
	"github.com/zakiyalmaya/online-store/model"
)

type Service interface {
	Create(request *model.CreateCartRequest) (*model.CartResponse, error)
	GetByParams(request *model.GetCartRequest) ([]*model.CartResponse, error)
	Delete(request *model.DeleteCartRequest) error
}