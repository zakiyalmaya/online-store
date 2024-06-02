package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/constant"
	"github.com/zakiyalmaya/online-store/model"
)

func AuthMiddleware(redcl *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Missing Authorization header"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the token signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(http.StatusUnauthorized, "Invalid signing method")
			}
			return []byte(constant.SecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Invalid or expired token"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Invalid token claims"))
		}

		username, ok := claims["username"].(string)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Invalid username in token claims"))
		}
		c.Locals("username", username)

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Invalid user_id in token claims"))
		}
		c.Locals("user_id", int(userID))

		// Check the token in Redis cache
		tokenCache, err := redcl.Get(context.Background(), constant.JWTPrefix+username).Result()
		if err != nil {
			log.Println("Redis error:", err)
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Invalid or expired token"))
		}

		if tokenCache != tokenString {
			return c.Status(http.StatusUnauthorized).JSON(model.HTTPErrorResponse("Invalid or expired token"))
		}

		return c.Next()
	}
}
