package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application"
	"github.com/zakiyalmaya/online-store/config"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/transport"
)

func main() {
	var sqliteDB, redisHost, redisPort, appPort string
	
	// Get environment variables
	sqliteDB = config.SQLITE_DB
	redisHost = config.REDIS_HOST
	redisPort = config.REDIS_PORT
	appPort = config.APP_PORT
	fmt.Println("configuration ", sqliteDB, redisHost, redisPort, appPort)

	// for testing run locally
	// sqliteDB = "./online_store.db"
	// redisHost = "localhost"
	// redisPort = "6379"
	// appPort = ":3000"

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
