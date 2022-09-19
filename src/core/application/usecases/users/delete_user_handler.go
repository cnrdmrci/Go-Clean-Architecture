package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/infastructure/persistence/mssql/mssql_connection"
	"go-clean-architecture/src/infastructure/persistence/mssql/repositories"
)

type DeleteUserByIdHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateDeleteUserHandler(logger interface_services.LoggerService) *DeleteUserByIdHandler {
	return &DeleteUserByIdHandler{
		Logger: logger,
	}
}

func (handler *DeleteUserByIdHandler) DeleteUserById(id int) error {
	db := mssql_connection.CreateConnection()
	tx := db.Begin()
	userRepository := repositories.CreateUserRepository(db)
	err := userRepository.DeleteUserByIdWithTransaction(tx, id)
	tx.Commit()
	return err
}
