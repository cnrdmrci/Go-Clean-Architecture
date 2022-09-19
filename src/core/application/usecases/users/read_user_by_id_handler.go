package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
	"go-clean-architecture/src/infastructure/persistence/mssql/mssql_connection"
	"go-clean-architecture/src/infastructure/persistence/mssql/repositories"
)

type ReadUserByIdHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateReadUserByIdHandler(logger interface_services.LoggerService) *ReadUserByIdHandler {
	return &ReadUserByIdHandler{
		Logger: logger,
	}
}

func (handler *ReadUserByIdHandler) ReadUserById(id int) (entities.User, error) {
	db := mssql_connection.CreateConnection()
	userRepository := repositories.CreateUserRepository(db)
	return userRepository.ReadUserById(id)
}
