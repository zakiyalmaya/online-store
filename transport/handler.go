package transport

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/middleware"
	"github.com/zakiyalmaya/online-store/transport/controller"
)

func Handler(application *application.Application, redcl *redis.Client, r *fiber.App) {
	ctrl := controller.NewController(application)

	r.Post("/customer", ctrl.Customer.Register)
	r.Post("/customer/login", ctrl.Customer.Login)
	r.Post("/customer/logout", middleware.AuthMiddleware(redcl), ctrl.Customer.Logout)
	
	r.Get("/categories", middleware.AuthMiddleware(redcl), ctrl.Category.GetAll)
	r.Post("/category", middleware.AuthMiddleware(redcl), ctrl.Category.Create)
}