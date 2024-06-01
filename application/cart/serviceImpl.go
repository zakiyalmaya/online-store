package cart

import (
	"database/sql"
	"fmt"
	"sync"

	cartEnum "github.com/zakiyalmaya/online-store/constant/cart"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/model"
)

type cartSvcImpl struct {
	repos *repository.Repositories
}

func NewCartService(repos *repository.Repositories) Service {
	return &cartSvcImpl{repos: repos}
}

func (c *cartSvcImpl) Create(request *model.CreateCartRequest) (*model.CartResponse, error) {
	// check product existence
	if err := c.checkProductExist(request); err != nil {
		return nil, err
	}

	// get existing cart
	// if cart not exist, create new cart
	// if cart exist, upsert product to cart
	activeCartStatus := int(cartEnum.CartStatusActive)
	cart, err := c.repos.Cart.GetByParams(&model.GetCartRequest{
		CustomerID: request.CustomerID,
		Status:     &activeCartStatus,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting active cart")
	}

	if len(cart) == 0 {
		cartEntity := request.ToEntity()
		newCart, err := c.repos.Cart.Create(cartEntity)
		if err != nil {
			return nil, fmt.Errorf("error creating cart")
		}
		cart = append(cart, newCart)
	} else {
		cartItems := make([]*model.CartItemEntity, len(request.Items))
		for i, item := range request.Items {
			cartItems[i] = &model.CartItemEntity{
				CartID:    cart[0].ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
			}
		}

		cart[0], err = c.repos.Cart.Upsert(cart[0].ID, cartItems)
		if err != nil {
			return nil, fmt.Errorf("error upserting cart")
		}
	}

	// there is only one active cart
	return cart[0].ToResponse(), nil
}

func (c *cartSvcImpl) checkProductExist(cart *model.CreateCartRequest) error {
	var wg sync.WaitGroup
	errorCh := make(chan error, len(cart.Items))

	for _, item := range cart.Items {
		wg.Add(1)

		go func(productID int) {
			defer wg.Done()

			_, err := c.repos.Product.GetByID(productID)
			if err != nil {
				if err == sql.ErrNoRows {
					errorCh <- fmt.Errorf("product not found: %d", productID)
					return
				}
				errorCh <- fmt.Errorf("error getting product by id: %d, %v", productID, err)
				return
			}
		}(item.ProductID)
	}

	wg.Wait()
	close(errorCh)

	for err := range errorCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *cartSvcImpl) GetByParams(request *model.GetCartRequest) ([]*model.CartResponse, error) {
	carts, err := c.repos.Cart.GetByParams(request)
	if err != nil {
		return nil, fmt.Errorf("error getting cart by params")
	}

	cartsResponse := make([]*model.CartResponse, len(carts))
	for i, cart := range carts {
		cartsResponse[i] = cart.ToResponse()
	}

	return cartsResponse, nil
}

func (c *cartSvcImpl) Delete(request *model.DeleteCartRequest) error {
	// check cart item existence
	cartItem, err := c.repos.Cart.GetItemByID(request.CartItemID)
	if err != nil {
		return fmt.Errorf("error getting cart item by id")
	}

	// delete the product from the cart if the cart status is active
	if err := c.repos.Cart.Delete(&model.DeleteCartRequest{
		ID:         cartItem.CartID,
		CartItemID: cartItem.ID,
		CustomerID: request.CustomerID,
		Status:     cartEnum.CartStatusActive,
	}); err != nil {
		return fmt.Errorf("error deleting cart item")
	}

	return nil
}
