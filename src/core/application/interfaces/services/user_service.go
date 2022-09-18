package interface_services

import "go-clean-architecture/src/core/domain/entities"

type UserService interface {
	CreateUser(user entities.User) (int64, error)
	UpdateUser(user entities.User) error
	ReadUserById(id int) (entities.User, error)
	DeleteUserById(id int) error
}
