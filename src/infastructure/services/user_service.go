package infrastructure_services

import (
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/application/usecases/users"
	"go-clean-architecture/src/core/domain/entities"
)

type UserService struct {
	Logger interface_services.LoggerService
}

func CreateUserService(logger interface_services.LoggerService) interface_services.UserService {
	return &UserService{
		Logger: logger,
	}
}

func (us *UserService) CreateUser(user entities.User) (int64, error) {
	createUserHandler := usecases.CreateCreateUserHandler(us.Logger)
	return createUserHandler.CreateUser(user)
}

func (us *UserService) UpdateUser(user entities.User) error {
	updateUserHandler := usecases.CreateUpdateUserHandler(us.Logger)
	return updateUserHandler.UpdateUser(user)
}

func (us *UserService) ReadUserById(id int) (entities.User, error) {
	readUserByIdHandler := usecases.CreateReadUserByIdHandler(us.Logger)
	return readUserByIdHandler.ReadUserById(id)
}

func (us *UserService) DeleteUserById(id int) error {
	deleteUserHandler := usecases.CreateDeleteUserHandler(us.Logger)
	return deleteUserHandler.DeleteUserById(id)
}
