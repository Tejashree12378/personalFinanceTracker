package limits

import (
	"context"
	"net/http"
	"strconv"

	"personalFinanceTracker/internal/app/controllers/models"
	serviceModel "personalFinanceTracker/internal/app/services/models"

	"github.com/gin-gonic/gin"
)

type limitService interface {
	Create(ctx context.Context, limit *serviceModel.Limit) error
	GetByID(ctx context.Context, id int) (*serviceModel.Limit, error)
	Update(ctx context.Context, limit *serviceModel.Limit) error
	Delete(ctx context.Context, id int) error
}

type LimitController struct {
	service limitService
}

func NewLimitController(s limitService) *LimitController {
	return &LimitController{service: s}
}

func (lc *LimitController) CreateLimit(c *gin.Context) {
	var req models.LimitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := lc.service.Create(c.Request.Context(), req.ToServiceModel()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create limit"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "limit created"})
}

func (lc *LimitController) GetLimit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	limit, err := lc.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "limit not found"})
		return
	}

	c.JSON(http.StatusOK, limit)
}

func (lc *LimitController) UpdateLimit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req models.LimitUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit := req.ToServiceModel()
	limit.ID = id

	if err := lc.service.Update(c.Request.Context(), limit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update limit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "limit updated"})
}

func (lc *LimitController) DeleteLimit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := lc.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete limit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "limit deleted"})
}
