package repositories

import "gorm.io/gorm"

type EmployeeRepository interface {
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}
