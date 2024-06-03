package customer

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/zakiyalmaya/online-store/constant"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	mockCustomerRepo "github.com/zakiyalmaya/online-store/mocks/infrastructure/repository/customer"
	"github.com/zakiyalmaya/online-store/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	mockCustomerRepository *mockCustomerRepo.MockRepository
	mockRedisClient        *redis.Client
	customerSvc            Service
)

func Setup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRedisServer, err := miniredis.Run()
	if err != nil {
		t.Fatalf(err.Error())
	}

	mockRedisClient = redis.NewClient(&redis.Options{
		Addr: mockRedisServer.Addr(),
	})

	mockCustomerRepository = mockCustomerRepo.NewMockRepository(mockCtl)
	customerSvc = NewCustomerService(&repository.Repositories{
		Customer: mockCustomerRepository,
		RedCl:    mockRedisClient,
	})
}

func TestRegister(t *testing.T) {
	Setup(t)

	request := &model.CustomerRequest{
		Name:     "name",
		Username: "username",
		Password: "password",
	}

	mockHashedPass := []byte("password")
	patched := gomonkey.ApplyFunc(bcrypt.GenerateFromPassword, func([]byte, int) ([]byte, error) {
		return mockHashedPass, nil
	})
	defer patched.Reset()

	testCases := []struct {
		name    string
		request *model.CustomerRequest
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given valid request when register then return success",
			request: request,
			mock: func() {
				mockCustomerRepository.EXPECT().Create(&model.CustomerEntity{
					Name:     request.Name,
					Username: request.Username,
					Password: string(mockHashedPass),
				}).Return(nil).Times(1)
			},
			wantErr: false,
		},
		{
			name:    "Given error create customer to db when register then return error",
			request: request,
			mock: func() {
				mockCustomerRepository.EXPECT().Create(&model.CustomerEntity{
					Name:     request.Name,
					Username: request.Username,
					Password: string(mockHashedPass),
				}).Return(errors.New("error")).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := customerSvc.Register(request)
			if (err != nil) != tc.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	Setup(t)

	request := &model.AuthRequest{
		Username: "username",
		Password: "password",
	}

	mockHashedPass := []byte("password")
	patched := gomonkey.ApplyFunc(bcrypt.CompareHashAndPassword, func([]byte, []byte) error {
		return nil
	})
	defer patched.Reset()

	testCases := []struct {
		name    string
		request *model.AuthRequest
		mock    func()
		wantErr bool
	}{
		{
			name:    "Given success when login then return success",
			request: request,
			mock: func() {
				mockCustomerRepository.EXPECT().GetByUsername(request.Username).Return(&model.CustomerEntity{
					Password: string(mockHashedPass),
				}, nil).Times(1)

				mockRedisClient.Set(context.Background(), constant.JWTPrefix+request.Username, gomock.Any(), constant.SessionExpire).Err()
			},
			wantErr: false,
		},
		{
			name:    "Given error when login then return error",
			request: request,
			mock: func() {
				mockCustomerRepository.EXPECT().GetByUsername(request.Username).Return(nil, errors.New("error")).Times(1)
			},
			wantErr: true,
		},
		{
			name:    "Given error user not found request when login then return error",
			request: request,
			mock: func() {
				mockCustomerRepository.EXPECT().GetByUsername(request.Username).Return(nil, sql.ErrNoRows).Times(1)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			_, err := customerSvc.Login(request)
			if (err != nil) != tc.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	Setup(t)

	testCases := []struct {
		name     string
		username string
		mock     func()
		wantErr  bool
	}{
		{
			name:     "Given success when logout then return success",
			username: "username",
			mock:     func() {
				mockRedisClient.Del(context.Background(), constant.JWTPrefix+"username").Err()
			},
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := customerSvc.Logout(tc.username)
			if (err != nil) != tc.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
