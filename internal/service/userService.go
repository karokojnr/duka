package service

import (
	"errors"
	"github.com/karokojnr/duka/internal/domain"
	"github.com/karokojnr/duka/internal/dto"
	"github.com/karokojnr/duka/internal/helper"
	"github.com/karokojnr/duka/internal/repository"
	"time"
)

type UserService struct {
	UsrRepo repository.UserRepository
	Auth    helper.Auth
}

func (svc UserService) Register(input dto.UserRegisterDto) (string, error) {

	hashedPassword, err := svc.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	usr, err := svc.UsrRepo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})
	return svc.Auth.GenerateToken(usr.ID, usr.Email, usr.UserRole)
}

func (svc UserService) Login(input dto.UserLoginDto) (string, error) {
	usr, err := svc.UsrRepo.GetUser(input.Email)

	if err != nil {
		return "", errors.New("no user found with the email provided")
	}

	err = svc.Auth.VerifyPassword(input.Password, usr.Password)
	if err != nil {
		return "", err
	}

	return svc.Auth.GenerateToken(usr.ID, usr.Email, usr.UserRole)
}

func (svc UserService) findUserByEmail(email string) (*domain.User, error) {
	usr, err := svc.UsrRepo.GetUser(email)
	return &usr, err
}

func (svc UserService) isVerified(id uint) bool {
	currentUser, err := svc.UsrRepo.GetUserById(id)
	return err == nil && currentUser.IsVerified
}

func (svc UserService) GetVerificationCode(usr domain.User) (int, error) {

	if svc.isVerified(usr.ID) {
		return 0, nil
	}

	code, err := svc.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	user := domain.User{
		ExpiryTime: time.Now().Add(30 * time.Minute),
		Code:       code,
	}

	_, err = svc.UsrRepo.UpdateUser(usr.ID, user)
	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	// todo: send sms

	return code, nil
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
