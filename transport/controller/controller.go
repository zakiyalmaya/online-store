package controller

import (
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/transport/controller/category"
)

type Controller struct {
	Category *category.Controller
}

func NewController(application *application.Application) *Controller {
	return &Controller{
		Category: category.NewCategoryController(application.CategorySvc),
	}
}