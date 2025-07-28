package repositories

import (
	"github.com/kryast/Crud-8.git/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *models.Employee) error
	GetAll() ([]models.Employee, error)
	GetByID(id uint) (*models.Employee, error)
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

func (r *employeeRepository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) GetByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.First(employee, id).Error
	return &employee, err
}
