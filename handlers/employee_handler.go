package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-8.git/models"
	"github.com/kryast/Crud-8.git/services"
)

type EmployeeHandler struct {
	service services.EmployeeService
}

func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	employees, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.ID = uint(id)
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee soft deleted"})
}
