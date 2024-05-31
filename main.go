package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/transport"
)

func main() {
	// instatiate repository
	db := repository.DBConnection()
	defer db.Close()

	repository := repository.NewRespository(db)

	// instantiate application
	application := application.NewApplication(repository)

	// instantiate fiber
	r := fiber.New()

	// instantiate transport
	transport.Handler(application, r)

	r.Listen(":3000")
}