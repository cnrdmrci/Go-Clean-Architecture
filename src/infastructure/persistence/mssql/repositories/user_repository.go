package repositories

import (
	_ "github.com/denisenkom/go-mssqldb"
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/domain/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) interface_repositories.UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (userRepository *UserRepository) CreateUser(user entities.User) (int64, error) {
	return userRepository.createUser(userRepository.Db, user)
}

func (userRepository *UserRepository) CreateUserWithTransaction(db *gorm.DB, user entities.User) (int64, error) {
	return userRepository.createUser(db, user)
}

func (userRepository *UserRepository) createUser(db *gorm.DB, user entities.User) (int64, error) {
	db.Create(&entities.User{Name: user.Name, Gender: user.Gender, Age: user.Age})
	return 0, nil
}

func (userRepository *UserRepository) UpdateUser(user entities.User) error {
	return userRepository.updateUser(userRepository.Db, user)
}

func (userRepository *UserRepository) UpdateUserWithTransaction(db *gorm.DB, user entities.User) error {
	return userRepository.updateUser(db, user)
}

func (userRepository *UserRepository) updateUser(db *gorm.DB, user entities.User) error {
	var tempUser entities.User
	userRepository.Db.Where("Id = ?", user.ID).First(&tempUser)

	tempUser.Name = user.Name
	tempUser.Gender = user.Gender
	tempUser.Age = user.Age

	userRepository.Db.Save(&tempUser)
	return nil
}

func (userRepository *UserRepository) ReadUserById(id int) (entities.User, error) {
	var user entities.User
	userRepository.Db.Where("Id = ?", id).First(&user)
	return user, nil
}

func (userRepository *UserRepository) DeleteUserById(id int) error {
	return userRepository.deleteUserById(userRepository.Db, id)
}

func (userRepository *UserRepository) DeleteUserByIdWithTransaction(db *gorm.DB, id int) error {
	return userRepository.deleteUserById(db, id)
}

func (userRepository *UserRepository) deleteUserById(db *gorm.DB, id int) error {
	userRepository.Db.Delete(&entities.User{}, id)
	return nil
}
