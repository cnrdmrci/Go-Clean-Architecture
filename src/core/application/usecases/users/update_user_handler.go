package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
)

type UpdateUserHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateUpdateUserHandler(logger interface_services.LoggerService, userRepository interface_repositories.UserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{
		Logger:         logger,
		UserRepository: userRepository,
	}
}

func (handler *UpdateUserHandler) UpdateUser(user entities.User) error {
	return handler.UserRepository.UpdateUser(user)
}
