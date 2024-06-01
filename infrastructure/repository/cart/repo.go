package cart

import "github.com/zakiyalmaya/online-store/model"

type Repository interface {
	Create(cart *model.CartEntity) (*model.CartEntity, error)
	GetByParams(request *model.GetCartRequest) ([]*model.CartEntity, error)
	Upsert(cartID int, items []*model.CartItemEntity) (*model.CartEntity, error)
	Delete(request *model.DeleteCartRequest) error
	GetItemByID(cartItemID int) (*model.CartItemEntity, error)
	GetByID(cartID int) (*model.CartEntity, error)
}
