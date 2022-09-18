package common

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "go-clean-architecture/docs"
	"go-clean-architecture/src/api/controllers"
	"go-clean-architecture/src/infastructure/services"
)

func CreateRoutes(router *gin.Engine) {
	logger := CreateLogger()
	userController := controllers.CreateUserController(infrastructure_services.CreateUserService(logger))

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", userController.CreateUser)
			users.PUT("/", userController.UpdateUser)
			users.GET("/:id", userController.ReadUserById)
			users.DELETE("/:id", userController.DeleteUserById)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
