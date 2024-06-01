package controller

import (
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/transport/controller/category"
	"github.com/zakiyalmaya/online-store/transport/controller/customer"
	"github.com/zakiyalmaya/online-store/transport/controller/product"
)

type Controller struct {
	Category *category.Controller
	Customer *customer.Controller
	Product  *product.Controller
}

func NewController(application *application.Application) *Controller {
	return &Controller{
		Category: category.NewCategoryController(application.CategorySvc),
		Customer: customer.NewCategoryController(application.CustomerSvc),
		Product:  product.NewProductController(application.ProductSvc),
	}
}
