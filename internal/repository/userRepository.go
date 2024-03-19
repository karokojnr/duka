package repository

import "os/user"

type UserRepository interface {
	CreateUser(usr user.User) (user.User, error)
	GetUser(email string) (user.User, error)
	GetUserById(id uint) (user.User, error)
	UpdateUser(id uint, usr user.User) (user.User, error)
}

type userRepository struct {
}
