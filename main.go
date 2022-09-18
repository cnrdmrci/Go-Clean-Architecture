package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/infastructure/services/common"
)

// @title Clean Architecture Swagger API
// @version 1.0
// @description User CRUD endpoints
// @termsOfService http://swagger.io/terms/
// @contact.name Caner Demirci
// @contact.email none
// @contact.url none
// @BasePath /
// @host localhost:8080
func main() {
	fmt.Println("Program started.")

	router := gin.Default()
	common.CreateRoutes(router)

	_ = router.Run(":8080")
}
