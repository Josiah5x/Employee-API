package repository

import (
	"github.com/ellipizle/golang-mvc/model"
)

type Repository interface {
	AddEmployee(employee *model.Employee) (*model.Employee, error)
	EditEmployee(employee *model.Employee) (*model.Employee, error)
	GetEmployee(id string) (*model.Employee, error)
	DeleteEmployee(id string) error
	GetAllEmployee() ([]*model.Employee, error)
}
