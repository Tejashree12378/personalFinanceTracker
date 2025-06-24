package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	expenseController "personalFinanceTracker/internal/app/controllers/expenses"
	limitController "personalFinanceTracker/internal/app/controllers/limits"
	userController "personalFinanceTracker/internal/app/controllers/users"
	expenseRepository "personalFinanceTracker/internal/app/repositories/expenses"
	limitRepository "personalFinanceTracker/internal/app/repositories/limits"
	userRepository "personalFinanceTracker/internal/app/repositories/users"
	expenseService "personalFinanceTracker/internal/app/services/expenses"
	limitService "personalFinanceTracker/internal/app/services/limits"
	userService "personalFinanceTracker/internal/app/services/users"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api/v1")

	// User
	userRepo := userRepository.New(db)
	userSrv := userService.NewUserService(userRepo)
	userCtrl := userController.NewUserController(userSrv)

	userRoutes := v1.Group("/users")
	{
		userRoutes.POST("", userCtrl.CreateUser)
		userRoutes.GET("/:id", userCtrl.GetUserByID)
		userRoutes.PATCH("/:id", userCtrl.UpdateUser)
		userRoutes.DELETE("/:id", userCtrl.DeleteUser)
	}

	// expenses
	expenseRepo := expenseRepository.New(db)
	expenseSrv := expenseService.NewExpenseService(expenseRepo)
	expenseCtrl := expenseController.NewExpenseController(expenseSrv)

	expenseRoutes := v1.Group("/expenses")
	{
		expenseRoutes.POST("", expenseCtrl.CreateExpense)
		expenseRoutes.GET("/:id", expenseCtrl.GetExpenseByID)
		expenseRoutes.PATCH("/:id", expenseCtrl.UpdateExpense)
		expenseRoutes.DELETE("/:id", expenseCtrl.DeleteExpense)
	}

	// limits
	limitRepo := limitRepository.New(db)
	limitSrv := limitService.NewLimitService(limitRepo)
	limitCtrl := limitController.NewLimitController(limitSrv)

	limitRoutes := v1.Group("/expenses")
	{
		limitRoutes.POST("", limitCtrl.CreateLimit)
		limitRoutes.GET("/:id", limitCtrl.GetLimit)
		limitRoutes.PATCH("/:id", limitCtrl.UpdateLimit)
		limitRoutes.DELETE("/:id", limitCtrl.DeleteLimit)
	}
}
