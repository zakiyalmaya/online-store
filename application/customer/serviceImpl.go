package customer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zakiyalmaya/online-store/config"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/model"
	"golang.org/x/crypto/bcrypt"
)

type customerSvcImpl struct {
	repos *repository.Repositories
}

func NewCustomerService(repos *repository.Repositories) Service {
	return &customerSvcImpl{repos: repos}
}

func (c *customerSvcImpl) Register(request *model.CustomerRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password")
	}

	if err := c.repos.Customer.Create(&model.CustomerEntity{
		Name:        request.Name,
		Username:    request.Username,
		Password:    string(hashedPassword),
		Email:       request.Email,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}); err != nil {
		return fmt.Errorf("error creating customer")
	}

	return nil
}

func (c *customerSvcImpl) Login(request *model.AuthRequest) (*model.AuthResponse, error) {
	response, err := c.repos.Customer.GetByUsername(request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("error getting customer by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password)); err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	duration := config.SESSION_EXPIRE
	expirationTime := time.Now().Add(duration)
	claims := &model.AuthClaims{
		UserID:   response.ID,
		Username: response.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return nil, fmt.Errorf("failed to create token")
	}

	err = c.repos.RedCl.Set(context.Background(), config.JWT_PREFIX+response.Username, tokenString, duration).Err()
	if err != nil {
		log.Println("Failed to store token in Redis:", err.Error())
		return nil, fmt.Errorf("failed to store token in Redis")
	}

	return &model.AuthResponse{
		Username: response.Username, 
		Name: response.Name, 
		Token: tokenString,
	}, nil
}

func (c *customerSvcImpl) Logout(username string) error {
	if err := c.repos.RedCl.Del(context.Background(), config.JWT_PREFIX+username).Err(); err != nil {
		log.Println("Failed to delete token from Redis:", err.Error())
		return fmt.Errorf("failed to delete token from Redis")
	}

	return nil
}