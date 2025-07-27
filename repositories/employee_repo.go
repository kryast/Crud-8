package repositories

import (
	"github.com/kryast/Crud-8.git/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}
