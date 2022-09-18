package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"go-clean-architecture/src/core/application/interfaces/repositories"
	"go-clean-architecture/src/core/domain/entities"
	"golang.org/x/net/context"
	"log"
)

var server = "server_address"
var port = 1433
var user = "user"
var password = "pass"

var db *sql.DB

type UserRepository struct{}

func CreateUserRepository() interface_repositories.UserRepository {
	return &UserRepository{}
}

func (userRepository *UserRepository) CreateUser(user entities.User) (int64, error) {
	createDatabaseConnection()
	defer db.Close()

	var err error
	ctx := context.Background()

	insertSql := "INSERT INTO TestSchema.Users (Name, Gender, Age) VALUES (@Name, @Gender, @Age); select convert(bigint, SCOPE_IDENTITY());"

	stmt, err := db.Prepare(insertSql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", user.Name),
		sql.Named("Gender", user.Gender),
		sql.Named("Age", user.Age),
	)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (userRepository *UserRepository) UpdateUser(user entities.User) error {
	createDatabaseConnection()
	defer db.Close()

	ctx := context.Background()

	tsql := fmt.Sprintf("UPDATE TestSchema.Users SET Name = @Name, Gender = @Gender, Age = @Age WHERE Id = @Id")

	_, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", user.Name),
		sql.Named("Gender", user.Gender),
		sql.Named("Age", user.Age),
		sql.Named("Id", user.ID))
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *UserRepository) ReadUserById(id int) (entities.User, error) {
	createDatabaseConnection()
	defer db.Close()

	ctx := context.Background()

	tsql := fmt.Sprintf("SELECT Id, Name, Gender, Age FROM TestSchema.Users WHERE Id = @Id;")

	rows, err := db.QueryContext(ctx, tsql, sql.Named("Id", id))
	if err != nil {
		return entities.User{}, err
	}
	defer rows.Close()

	var name, gender string
	var age int

	rows.Next()
	scanErr := rows.Scan(&id, &name, &gender, &age)
	if scanErr != nil {
		return entities.User{}, scanErr
	}

	entity := entities.User{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
	}

	return entity, nil

}

func (userRepository *UserRepository) DeleteUserById(id int) error {
	createDatabaseConnection()
	defer db.Close()

	ctx := context.Background()

	tsql := fmt.Sprintf("DELETE FROM TestSchema.Users WHERE Id = @Id;")

	_, err := db.ExecContext(ctx, tsql, sql.Named("Id", id))
	if err != nil {
		return err
	}

	return nil
}

func createDatabaseConnection() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
}
