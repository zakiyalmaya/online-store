package application

import (
	"github.com/zakiyalmaya/online-store/application/cart"
	"github.com/zakiyalmaya/online-store/application/category"
	"github.com/zakiyalmaya/online-store/application/customer"
	"github.com/zakiyalmaya/online-store/application/product"
	"github.com/zakiyalmaya/online-store/application/transaction"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
)

type Application struct {
	CategorySvc    category.Service
	CustomerSvc    customer.Service
	ProductSvc     product.Service
	CartSvc        cart.Service
	TransactionSvc transaction.Service
}

func NewApplication(repos *repository.Repositories) *Application {
	return &Application{
		CategorySvc:    category.NewCategoryService(repos),
		CustomerSvc:    customer.NewCustomerService(repos),
		ProductSvc:     product.NewProductService(repos),
		CartSvc:        cart.NewCartService(repos),
		TransactionSvc: transaction.NewTransactionService(repos),
	}
}
