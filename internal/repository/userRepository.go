package repository

import (
	"github.com/karokojnr/duka/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	GetUser(email string) (domain.User, error)
	GetUserById(id uint) (domain.User, error)
	UpdateUser(id uint, usr domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (r userRepository) GetUser(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (r userRepository) GetUserById(id uint) (domain.User, error) {
	return domain.User{}, nil
}

func (r userRepository) UpdateUser(id uint, usr domain.User) (domain.User, error) {
	return domain.User{}, nil
}
