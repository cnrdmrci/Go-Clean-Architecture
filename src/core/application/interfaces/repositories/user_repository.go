package interface_repositories

import "go-clean-architecture/src/core/domain/entities"

type UserRepository interface {
	CreateUser(user entities.User) (int64, error)
	UpdateUser(user entities.User) error
	ReadUserById(id int) (entities.User, error)
	DeleteUserById(id int) error
}
