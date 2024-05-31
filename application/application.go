package application

import (
	"github.com/zakiyalmaya/online-store/application/category"
	"github.com/zakiyalmaya/online-store/application/customer"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
)

type Application struct {
	CategorySvc category.Service
	CustomerSvc customer.Service
}

func NewApplication(repos *repository.Repositories) *Application {
	return &Application{
		CategorySvc: category.NewCategoryService(repos),
		CustomerSvc: customer.NewCustomerService(repos),
	}
}
