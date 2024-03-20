package repository

import (
	"errors"
	"github.com/karokojnr/duka/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
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

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {

	err := r.db.Create(&usr).Error
	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create new user")
	}
	return usr, nil
}

func (r userRepository) GetUser(email string) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("get user error %v", err)
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r userRepository) GetUserById(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "id=?", id).Error
	if err != nil {
		log.Printf("get user error %v", err)
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r userRepository) UpdateUser(id uint, usr domain.User) (domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(usr).Error
	if err != nil {
		log.Printf("update user error %v", err)
		return domain.User{}, errors.New("failed to user")
	}
	return user, nil
}
