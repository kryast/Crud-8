package services

import (
	"github.com/kryast/Crud-8.git/models"
	"github.com/kryast/Crud-8.git/repositories"
)

type EmployeeService interface {
	Create(employee *models.Employee) error
}

type employeeService struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s *employeeService) Create(employee *models.Employee) error {
	return s.repo.Create(employee)
}
