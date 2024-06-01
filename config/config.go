package config

import "time"

const (
	APP_PORT = ":3000"

	DATABASE_NAME = "./online_store.db"

	REDIS_HOST = "localhost"
	REDIS_PORT = "6379"
	REDIS_PASS = ""

	SECRET_KEY = "online-store-secret"

	SESSION_EXPIRE = 10 * time.Minute

	JWT_PREFIX = "jwt-token-"
)