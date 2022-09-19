package mssql_connection

import (
	"fmt"
	"go-clean-architecture/src/infastructure/common"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func CreateConnection() *gorm.DB {
	config, configErr := infrastructure_common.LoadConfig()
	if configErr != nil {
		fmt.Println(configErr)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		config.Db.Address, config.Db.User, config.Db.Password, config.Db.Port)

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.Db.Schema + ".",
		},
	})
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}

	return db
}
