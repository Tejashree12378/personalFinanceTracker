package expenses

import (
	"context"
	"net/http"
	"strconv"

	"personalFinanceTracker/internal/app/controllers/models"

	serviceModel "personalFinanceTracker/internal/app/services/models"

	"github.com/gin-gonic/gin"
)

type expenseService interface {
	CreateExpense(ctx context.Context, expense *serviceModel.Expense) error
	UpdateExpense(ctx context.Context, expense *serviceModel.Expense) error
	GetExpenseByID(ctx context.Context, id int) (*serviceModel.Expense, error)
	DeleteExpense(ctx context.Context, id int) error
}

type ExpenseController struct {
	svc expenseService
}

func NewExpenseController(svc expenseService) *ExpenseController {
	return &ExpenseController{svc: svc}
}

func (ec *ExpenseController) CreateExpense(c *gin.Context) {
	var req *models.ExpenseCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := req.ToServiceModel()
	if err := ec.svc.CreateExpense(c.Request.Context(), expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create expense"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

func (ec *ExpenseController) UpdateExpense(c *gin.Context) {
	id := c.Param("id")

	expenseID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req *models.ExpenseUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = expenseID

	if err := ec.svc.UpdateExpense(c.Request.Context(), req.ToServiceModel()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})
}

func (ec *ExpenseController) GetExpenseByID(c *gin.Context) {
	id := c.Param("id")
	expenseID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	expense, err := ec.svc.GetExpenseByID(c.Request.Context(), expenseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (ec *ExpenseController) DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	expenseID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := ec.svc.DeleteExpense(c.Request.Context(), expenseID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete expense"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
