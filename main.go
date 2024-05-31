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
	redcl := repository.RedisClient()
	defer db.Close()

	repository := repository.NewRespository(db, redcl)

	// instantiate application
	application := application.NewApplication(repository)

	// instantiate fiber
	r := fiber.New()

	// instantiate transport
	transport.Handler(application, redcl, r)

	r.Listen(":3000")
}