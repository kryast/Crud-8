package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-8.git/handlers"
	"github.com/kryast/Crud-8.git/repositories"
	"github.com/kryast/Crud-8.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewEmployeeRepository(db)
	svc := services.NewEmployeeService(repo)
	h := handlers.NewEmployeeHandler(svc)

	r.POST("/employees", h.Create)
	r.GET("/employees", h.GetAll)
	r.GET("/employees/:id", h.GetByID)

	return r
}
