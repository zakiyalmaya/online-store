package constant

import "time"

const (
	SecretKey     = "online-store-secret"
	SessionExpire = 10 * time.Minute
	JWTPrefix     = "jwt-token-"

	DefaultLimit = 10
	DefaultPage  = 1
)