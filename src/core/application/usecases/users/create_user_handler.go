package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
	"go-clean-architecture/src/infastructure/persistence/mssql/mssql_connection"
	"go-clean-architecture/src/infastructure/persistence/mssql/repositories"
)

type CreateUserHandler struct {
	Logger interface_services.LoggerService
}

func CreateCreateUserHandler(logger interface_services.LoggerService) *CreateUserHandler {
	return &CreateUserHandler{
		Logger: logger,
	}
}

func (handler *CreateUserHandler) CreateUser(user entities.User) (int64, error) {
	db := mssql_connection.CreateConnection()
	tx := db.Begin()
	userRepository := repositories.CreateUserRepository(db)
	id, err := userRepository.CreateUserWithTransaction(tx, user)
	tx.Commit()

	return id, err
}
