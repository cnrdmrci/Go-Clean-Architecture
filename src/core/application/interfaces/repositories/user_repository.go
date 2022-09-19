package interface_repositories

import (
	"go-clean-architecture/src/core/domain/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserWithTransaction(db *gorm.DB, user entities.User) (int64, error)
	CreateUser(user entities.User) (int64, error)
	UpdateUser(user entities.User) error
	UpdateUserWithTransaction(db *gorm.DB, user entities.User) error
	ReadUserById(id int) (entities.User, error)
	DeleteUserById(id int) error
	DeleteUserByIdWithTransaction(db *gorm.DB, id int) error
}
