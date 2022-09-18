package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
)

type ReadUserByIdHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateReadUserByIdHandler(logger interface_services.LoggerService, userRepository interface_repositories.UserRepository) *ReadUserByIdHandler {
	return &ReadUserByIdHandler{
		Logger:         logger,
		UserRepository: userRepository,
	}
}

func (handler *ReadUserByIdHandler) ReadUserById(id int) (entities.User, error) {
	return handler.UserRepository.ReadUserById(id)
}
