package services

import "github.com/kryast/Crud-8.git/repositories"

type EmployeeService interface {
}

type employeeService struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}
