package service

import (
	"github.com/karokojnr/duka/internal/domain"
	"github.com/karokojnr/duka/internal/dto"
	"log"
)

type UserService struct {
	// todo
}

func (svc UserService) Register(input dto.UserRegisterDto) (string, error) {
	log.Println(input)
	// todo: handle user registration
	return "user-token", nil
}

func (svc UserService) Login(input any) (string, error) {
	return "", nil
}

func (svc UserService) findUserByEmail(email string) (*domain.User, error) {
	// todo: perform some db operation
	// todo: business logic
	return nil, nil
}

func (svc UserService) GetVerificationCode(usr domain.User) (int, error) {
	return 0, nil
}

func (svc UserService) VerifyCode(userId uint, code int) error {
	return nil
}

func (svc UserService) CreateProfile(userId uint, input any) error {
	return nil
}

func (svc UserService) GetProfile(userId uint) (*domain.User, error) {
	return nil, nil
}

func (svc UserService) UpdateProfile(userId uint, input any) error {
	return nil
}

func (svc UserService) AddRole(userId uint, input any) (string, error) {
	return "", nil
}

func (svc UserService) GetCart(userId uint) ([]interface{}, error) {
	return nil, nil
}

func (svc UserService) AddToCart(input any, usr domain.User) ([]interface{}, error) {
	return nil, nil
}

func (svc UserService) CreateOrder(usr domain.User) (int, error) {
	return 0, nil
}

func (svc UserService) GetOrders(usr domain.User) ([]interface{}, error) {
	return nil, nil
}

func (svc UserService) GetOrderById(orderId, userId uint) (interface{}, error) {
	return nil, nil
}
