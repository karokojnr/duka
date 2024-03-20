package service

import (
	"errors"
	"fmt"
	"github.com/karokojnr/duka/internal/domain"
	"github.com/karokojnr/duka/internal/dto"
	"github.com/karokojnr/duka/internal/repository"
)

type UserService struct {
	UsrRepo repository.UserRepository
}

func (svc UserService) Register(input dto.UserRegisterDto) (string, error) {
	usr, err := svc.UsrRepo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})
	// todo: generate token
	userInfo := fmt.Sprintf("%v, %v, %v", usr.ID, usr.Email, usr.UserRole)
	return userInfo, err
}

func (svc UserService) Login(input dto.UserLoginDto) (string, error) {
	usr, err := svc.UsrRepo.GetUser(input.Email)
	// todo:compare password and generate token
	if err != nil {
		return "", errors.New("no user found with the email provided")
	}
	return usr.Email, err
}

func (svc UserService) findUserByEmail(email string) (*domain.User, error) {
	usr, err := svc.UsrRepo.GetUser(email)
	return &usr, err
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
