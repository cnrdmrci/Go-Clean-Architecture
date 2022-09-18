package controllers

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/src/core/application/interfaces/services"
	"go-clean-architecture/src/core/domain/entities"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService interface_services.UserService
}

func CreateUserController(userService interface_services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// CreateUser godoc
// @Summary create user.
// @Produce json
// @Accept json
// @Param json body entities.User true "User data"
// @Success 200 {string} string "user created."
// @Failure 500
// @Router /v1/users/ [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uc.UserService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "User created.", "id": id})
}

// ReadUserById godoc
// @Summary read user by id.
// @Produce json
// @Param   id   path int  true  "User ID"
// @Success 200 {string} user info.
// @Failure 404 {string} not found.
// @Router /v1/users/{id} [get]
func (uc *UserController) ReadUserById(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	user, err := uc.UserService.ReadUserById(aid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary update user
// @Produce json
// @Accept json
// @Param json body entities.User string "User data"
// @Success 200 {string} string "user updated."
// @Failure 500
// @Router /v1/users/ [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "User updated."})
}

// DeleteUserById godoc
// @Summary delete user.
// @Produce json
// @Accept json
// @Param   id   path int  true  "User ID"
// @Success 200 {string} string "user deleted."
// @Failure 500
// @Router /v1/users/{id} [delete]
func (uc *UserController) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	aid, _ := strconv.Atoi(id)
	err := uc.UserService.DeleteUserById(aid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "User deleted."})
}
