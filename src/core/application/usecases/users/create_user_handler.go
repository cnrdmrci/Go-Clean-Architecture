package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
)

type CreateUserHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateCreateUserHandler(logger interface_services.LoggerService, userRepository interface_repositories.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		Logger:         logger,
		UserRepository: userRepository,
	}
}

func (handler *CreateUserHandler) CreateUser(user entities.User) (int64, error) {
	return handler.UserRepository.CreateUser(user)
}
