package controller

import (
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/transport/controller/category"
	"github.com/zakiyalmaya/online-store/transport/controller/customer"
)

type Controller struct {
	Category *category.Controller
	Customer *customer.Controller
}

func NewController(application *application.Application) *Controller {
	return &Controller{
		Category: category.NewCategoryController(application.CategorySvc),
		Customer: customer.NewCategoryController(application.CustomerSvc),
	}
}
