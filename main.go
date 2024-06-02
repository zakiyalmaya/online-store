package main

import (
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/transport"
)

func main() {
	// Get environment variables
	sqliteDB := os.Getenv("SQLITE_DB")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	appPort := os.Getenv("APP_PORT")

	// instatiate repository
	db := repository.DBConnection(sqliteDB)
	redcl := repository.RedisClient(redisHost, redisPort)
	defer db.Close()

	repository := repository.NewRepository(db, redcl)

	// instantiate application
	application := application.NewApplication(repository)

	// instantiate fiber
	r := fiber.New()

	// instantiate transport
	transport.Handler(application, redcl, r)

	fmt.Println("Server is running on port", appPort)
	r.Listen(appPort)
}
