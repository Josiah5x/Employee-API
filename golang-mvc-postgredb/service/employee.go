package service

import (
	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo}
}

// GetEmployee returned employee of a given ID
func (s *Service) GetEmployee(id string) (*model.Employee, error) {
	return s.repo.GetEmployee(id)
}

// GetAllEmployee all employee
func (s *Service) GetAllEmployee() ([]*model.Employee, error) {
	return s.repo.GetAllEmployee()
}

// DeleteEmployee remove an employee of given id
func (s *Service) DeleteEmployee(id string) error {
	return s.repo.DeleteEmployee(id)
}

// GetEmployee returned employee of a given ID
func (s *Service) AddEmployee(emp *model.Employee) (*model.Employee, error) {
	return s.repo.AddEmployee(emp)
}

// EditEmployee update an employee record
func (s *Service) EditEmployee(emp *model.Employee) (*model.Employee, error) {
	return s.repo.EditEmployee(emp)
}
