package users

import (
	"context"
	"net/http"
	"strconv"

	"personalFinanceTracker/internal/app/controllers/models"
	serviceModel "personalFinanceTracker/internal/app/services/models"

	"github.com/gin-gonic/gin"
)

type userService interface {
	CreateUser(ctx context.Context, user *serviceModel.User) error
	GetUserByID(ctx context.Context, id int) (*serviceModel.User, error)
	UpdateUser(ctx context.Context, user *serviceModel.User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserController struct {
	service userService
}

func NewUserController(s userService) *UserController {
	return &UserController{s}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user *models.UserCreateRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateUser(ctx.Request.Context(), user.ToServiceModel()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	user, err := c.service.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user *models.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	user.ID = uint(id)
	if err := c.service.UpdateUser(ctx.Request.Context(), user.ToServiceModel()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := c.service.DeleteUser(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
