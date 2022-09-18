package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
)

type DeleteUserByIdHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateDeleteUserHandler(logger interface_services.LoggerService, userRepository interface_repositories.UserRepository) *DeleteUserByIdHandler {
	return &DeleteUserByIdHandler{
		Logger:         logger,
		UserRepository: userRepository,
	}
}

func (handler *DeleteUserByIdHandler) DeleteUserById(id int) error {
	return handler.UserRepository.DeleteUserById(id)
}
