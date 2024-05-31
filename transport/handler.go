package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/transport/controller"
)

func Handler(application *application.Application, r *fiber.App) {
	ctrl := controller.NewController(application)

	r.Get("/categories", ctrl.Category.GetAll)
	r.Post("/category", ctrl.Category.Create)
}