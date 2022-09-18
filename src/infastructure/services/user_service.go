package infrastructure_services

import (
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/application/usecases/users"
	"go-clean-architecture/src/core/domain/entities"
	"go-clean-architecture/src/infastructure/persistence/repositories"
)

type UserService struct {
	Logger         interface_services.LoggerService
	UserRepository interface_repositories.UserRepository
}

func CreateUserService(logger interface_services.LoggerService) interface_services.UserService {
	return &UserService{
		Logger:         logger,
		UserRepository: repositories.CreateUserRepository(),
	}
}

func (us *UserService) CreateUser(user entities.User) (int64, error) {
	createUserHandler := usecases.CreateCreateUserHandler(us.Logger, us.UserRepository)
	return createUserHandler.CreateUser(user)
}

func (us *UserService) UpdateUser(user entities.User) error {
	updateUserHandler := usecases.CreateUpdateUserHandler(us.Logger, us.UserRepository)
	return updateUserHandler.UpdateUser(user)
}

func (us *UserService) ReadUserById(id int) (entities.User, error) {
	readUserByIdHandler := usecases.CreateReadUserByIdHandler(us.Logger, us.UserRepository)
	return readUserByIdHandler.ReadUserById(id)
}

func (us *UserService) DeleteUserById(id int) error {
	deleteUserHandler := usecases.CreateDeleteUserHandler(us.Logger, us.UserRepository)
	return deleteUserHandler.DeleteUserById(id)
}
