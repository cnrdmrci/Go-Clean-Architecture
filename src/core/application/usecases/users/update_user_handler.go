package usecases

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
	"go-clean-architecture/src/infastructure/persistence/mssql/mssql_connection"
	"go-clean-architecture/src/infastructure/persistence/mssql/repositories"
)

type UpdateUserHandler struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateUpdateUserHandler(logger interface_services.LoggerService) *UpdateUserHandler {
	return &UpdateUserHandler{
		Logger: logger,
	}
}

func (handler *UpdateUserHandler) UpdateUser(user entities.User) error {
	db := mssql_connection.CreateConnection()
	tx := db.Begin()
	userRepository := repositories.CreateUserRepository(db)
	err := userRepository.UpdateUserWithTransaction(tx, user)
	tx.Commit()

	return err
}
